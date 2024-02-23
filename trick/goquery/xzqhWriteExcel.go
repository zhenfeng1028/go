package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
)

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
	// Request the HTML page.
	res, err := http.Get("http://xzqh.mca.gov.cn/defaultQuery?shengji=%BA%D3%B1%B1%CA%A1%A3%A8%BC%BD%A3%A9&diji=-1&xianji=-1")
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
				fmt.Printf("region:%s\t", region)
			case 1:
				resident := enc.ConvertString(td.Text())
				x.Resident = resident
				fmt.Printf("resident:%s\t", resident)
			case 2:
				population := td.Text()
				// 去除空格
				population = strings.Replace(population, " ", "", -1)
				// 去除换行符
				population = strings.Replace(population, "\n", "", -1)
				// 去除转义字符\t
				population = strings.Replace(population, "\t", "", -1)
				x.Population = population
				fmt.Printf("population:%s\t", population)
			case 3:
				area := td.Text()
				// 去除空格
				area = strings.Replace(area, " ", "", -1)
				// 去除换行符
				area = strings.Replace(area, "\n", "", -1)
				// 去除转义字符\t
				area = strings.Replace(area, "\t", "", -1)
				x.Area = area
				fmt.Printf("area:%s\t", area)
			case 4:
				x.Code = td.Text()
				fmt.Printf("code:%s\t", x.Code)
			case 5:
				x.AreaCode = td.Text()
				fmt.Printf("area_code:%s\t", x.AreaCode)
			case 6:
				x.PostCode = td.Text()
				fmt.Println("post_code:", x.PostCode)
			}
		})
		arrayXzqh = append(arrayXzqh, x)
	})
	// fmt.Println(arrayXzqh)
}

func writeExcel(arr []Xzqh) {
	xlsxNew := excelize.NewFile()
	newSheet := "Sheet1"
	xlsxNew.NewSheet(newSheet)

	for i, xzqh := range arr {
		t := reflect.TypeOf(xzqh)
		v := reflect.ValueOf(xzqh)
		for j := 0; j < t.NumField(); j++ {
			// f := t.Field(j)
			val := v.Field(j).Interface()
			zuobiao := ChangIndexToAxis(i, j)
			ModifyExcelCellByAxis(xlsxNew, newSheet, zuobiao, val)
		}
	}
	xlsxNew.SaveAs("./河北省.xlsx")
}

func main() {
	postScrape()
	writeExcel(arrayXzqh)
}

// 数组下标转换成excel坐标
func ChangIndexToAxis(intIndexX int, intIndexY int) string {
	arr := [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
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
