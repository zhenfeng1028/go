package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

// 这个只是一个简单的版本只是获取QQ邮箱并且没有进行封装操作，另外爬出来的数据也没有进行去重操作
var (
	reQQEmail = `(\d+)@qq\.com`
)

// 爬邮箱
func GetEmail() {
	// 1.去网站拿数据
	resp, err := http.Get("https://tieba.baidu.com/p/8299418946")
	HandleError(err, "http.Get")
	defer resp.Body.Close()
	// 2.读取页面内容
	pageBytes, err := io.ReadAll(resp.Body)
	HandleError(err, "io.ReadAll")
	// 字节转字符串
	pageStr := string(pageBytes)
	// 3.过滤数据，筛出qq邮箱
	re := regexp.MustCompile(reQQEmail)
	// -1代表取全部
	results := re.FindAllStringSubmatch(pageStr, -1)

	// 遍历结果
	for _, result := range results {
		fmt.Println("email:", result[0])
		fmt.Println("qq:", result[1])
	}
}

// 处理异常
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}
func main() {
	GetEmail()
}
