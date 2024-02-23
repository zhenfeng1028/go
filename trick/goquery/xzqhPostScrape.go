package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

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
	res, err := http.Get("http://xzqh.mca.gov.cn/defaultQuery?shengji=%B1%B1%BE%A9%CA%D0%A3%A8%BE%A9%A3%A9&diji=-1&xianji=-1")
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
	fmt.Println(arrayXzqh)
}

func main() {
	postScrape()
}
