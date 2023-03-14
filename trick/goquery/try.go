package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

var provinces = []string{
	"北京市（京）",
	"天津市（津）",
	"河北省（冀）",
	"山西省（晋）",
	"内蒙古自治区（内蒙古）",
	"辽宁省（辽）",
	"吉林省（吉）",
	"黑龙江省（黑）",
	"上海市（沪）",
	"江苏省（苏）",
	"浙江省（浙）",
	"安徽省（皖）",
	"福建省（闽）",
	"江西省（赣）",
	"山东省（鲁）",
	"河南省（豫）",
	"湖北省（鄂）",
	"湖南省（湘）",
	"广东省（粤）",
	"广西壮族自治区（桂）",
	"海南省（琼）",
	"重庆市（渝）",
	"四川省（川、蜀）",
	"贵州省（黔、贵）",
	"云南省（滇、云）",
	"西藏自治区（藏）",
	"陕西省（陕、秦）",
	"甘肃省（甘、陇）",
	"青海省（青）",
	"宁夏回族自治区（宁）",
	"新疆维吾尔自治区（新）",
}

type Xzqh struct {
	Region     string
	Resident   string
	Population string
	Area       string
	Code       string
	AreaCode   string
	PostCode   string
}

var arrayXzqh []Xzqh

func postScrape() {
	for _, province := range provinces {
		decodeBytes, _ := Utf8ToGbk([]byte(province))
		decodeStr := make([]string, 0)
		for _, decodeByte := range decodeBytes {
			valueHex := fmt.Sprintf("%X", decodeByte)
			decodeStr = append(decodeStr, valueHex)
		}
		reqStr := strings.Join(decodeStr, "%")
		reqStr = "%" + reqStr

		// Request the HTML page.
		url := fmt.Sprintf("http://xzqh.mca.gov.cn/defaultQuery?shengji=%s&diji=-1&xianji=-1", reqStr)
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}

		// Load the HTML document
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		var enc mahonia.Decoder
		enc = mahonia.NewDecoder("gbk")

		doc.Find("table.info_table tbody tr").Each(func(i int, tr *goquery.Selection) {

			x := Xzqh{}

			tr.Find("td").Each(func(ix int, td *goquery.Selection) {
				switch ix {
				case 0:
					region := enc.ConvertString(td.Text())
					// 去除空格
					region = strings.Replace(region, " ", "", -1)
					// 去除加号
					region = strings.Replace(region, "+", "", -1)
					x.Region = region
				case 1:
					resident := enc.ConvertString(td.Text())
					x.Resident = resident
				case 2:
					population := td.Text()
					// 去除空格
					population = strings.Replace(population, " ", "", -1)
					// 去除换行符
					population = strings.Replace(population, "\n", "", -1)
					// 去除转义字符\t
					population = strings.Replace(population, "\t", "", -1)
					x.Population = population
				case 3:
					area := td.Text()
					// 去除空格
					area = strings.Replace(area, " ", "", -1)
					// 去除换行符
					area = strings.Replace(area, "\n", "", -1)
					// 去除转义字符\t
					area = strings.Replace(area, "\t", "", -1)
					x.Area = area
				case 4:
					x.Code = td.Text()
				case 5:
					x.AreaCode = td.Text()
				case 6:
					x.PostCode = td.Text()
				}
			})
			arrayXzqh = append(arrayXzqh, x)
		})
		// fmt.Println(arrayXzqh)
	}
}

func main() {
	postScrape()
	writeExcel(arrayXzqh)
}

func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func writeExcel(arr []Xzqh) {
	xlsxNew := excelize.NewFile()
	newSheet := "Sheet1"
	xlsxNew.NewSheet(newSheet)

	for i, xzqh := range arr {
		t := reflect.TypeOf(xzqh)
		v := reflect.ValueOf(xzqh)
		for j := 0; j < t.NumField(); j++ {
			val := v.Field(j).Interface()
			zuobiao := ChangIndexToAxis(i, j)
			ModifyExcelCellByAxis(xlsxNew, newSheet, zuobiao, val)
		}
	}
	xlsxNew.SaveAs("./全国行政区划代码.xlsx")
}

// 数组下标转换成excel坐标
func ChangIndexToAxis(intIndexX int, intIndexY int) string {
	var arr = [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	intIndexY = intIndexY + 1
	resultY := ""
	for true {
		if intIndexY <= 26 {
			resultY = resultY + arr[intIndexY-1]
			break
		}
		mo := intIndexY % 26
		resultY = arr[mo-1] + resultY
		shang := intIndexY / 26
		if shang <= 26 {
			resultY = arr[shang-1] + resultY
			break
		}
		intIndexY = shang
	}
	return resultY + strconv.Itoa(intIndexX+1)
}

// 修改excel表格里的值
func ModifyExcelCellByAxis(xlsx *excelize.File, sheet string, axis string, value interface{}) int {
	xlsx.SetCellValue(sheet, axis, value)
	return 0
}
