package main

import (
	"bytes"
	"crypto/md5"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type BrokerImage struct {
	BrokerName   string `json:"BrokerName"`
	BrokerMobile string `json:"BrokerMobile"`
	URL          string `json:"Url"`
}

var urlFormat = "https://zp.wodedagong.com/?r_str=ph711h1iaa6hcn00&broker_name=%s&come=timeline&hiredate=1614556800000&mobile=%s&sharedate=2025-03-03&sign=%s&utm_source=jjr&utm_medium=pyq&article_id=124"

func main() {
	// 打开 CSV 文件
	file, err := os.Open("3-13.csv")
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return
	}
	defer file.Close()

	// 创建 CSV Reader
	reader := csv.NewReader(file)

	// 读取所有记录
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("读取 CSV 文件失败:", err)
		return
	}

	bis := make([]BrokerImage, 0, 146)
	for _, record := range records {
		h := md5.New()
		io.WriteString(h, "wodedagong")
		io.WriteString(h, record[0])
		io.WriteString(h, record[1])
		io.WriteString(h, "1614556800000")
		sign := fmt.Sprintf("%x", h.Sum(nil))
		url := fmt.Sprintf(urlFormat, record[0], record[1], sign)
		// fmt.Println(url)
		bi := BrokerImage{
			BrokerName:   record[0],
			BrokerMobile: record[1],
			URL:          url,
		}
		bis = append(bis, bi)
	}
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	if err := jsonEncoder.Encode(bis); err != nil {
		panic(err)
	}

	fmt.Println(bf.String())
}
