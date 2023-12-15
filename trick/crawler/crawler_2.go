package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

var (
	// w代表大小写字母+数字+下划线
	reEmail = `\w+@\w+\.\w+`
	// s?有或者没有s
	// +代表出现1次或多次
	// \s\S各种字符
	// +?代表贪婪模式
	reLinke = `href="(https?://[\s\S]+?)"`
	rePhone = `1[3456789]\d\s?\d{4}\s?\d{4}`
	reImg   = `https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
)

// 处理异常
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}

// 抽取根据url获取内容
func GetPageStr(url string) (pageStr string) {
	resp, err := http.Get(url)
	HandleError(err, "http.Get url")
	defer resp.Body.Close()
	// 2.读取页面内容
	pageBytes, err := io.ReadAll(resp.Body)
	HandleError(err, "io.ReadAll")
	// 字节转字符串
	pageStr = string(pageBytes)
	return pageStr
}

func main() {
	// 2.抽取的爬邮箱
	// GetEmail2("https://tieba.baidu.com/p/8299418946")
	// 3.爬链接
	// GetLink("https://www.baidu.com/s?wd=%E5%8C%97%E4%BA%AC%E4%BD%8F%E6%88%BF%E5%85%AC%E7%A7%AF%E9%87%91%E6%89%A7%E8%A1%8C%E6%96%B0%E6%94%BF&tn=baidutop10&rsv_idx=2&usm=1&ie=utf-8&rsv_pq=fbfdf9710032b5bb&oq=AI%E7%94%A8%E8%8D%AF%E8%AF%B4%E6%98%8E%E4%B9%A6&rsv_t=966aARmwbiN1gMMB6seVwimtp2E7zxserKj8PfZw%2FgKWhYAzGEaM0oNe%2FKmcLDgIsg&rqid=fbfdf9710032b5bb&rsf=8cdaa38fc9fea91e45a8d70e06978ecc_1_15_15&rsv_dl=0_right_fyb_pchot_20811&sa=0_right_fyb_pchot_20811")
	// 4.爬手机号
	// GetPhone("https://www.zhaohaowang.com/")
	// 5.爬图片
	GetImg("http://image.baidu.com/search/index?tn=baiduimage&ps=1&ct=201326592&lm=-1&cl=2&nc=1&ie=utf-8&word=%E7%BE%8E%E5%A5%B3")
}

func GetEmail2(url string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reEmail)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println(result)
	}
}

// 爬链接
func GetLink(url string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reLinke)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println(result[1])
	}
}

// 爬手机号
func GetPhone(url string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(rePhone)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println(result)
	}
}

func GetImg(url string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reImg)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println(result[0])
	}
}
