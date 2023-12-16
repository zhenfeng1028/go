package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

// 并发爬思路：
// 1.初始化数据管道
// 2.爬虫写出：26个协程向管道中添加图片链接
// 3.任务统计协程：检查26个任务是否都完成，完成则关闭数据管道
// 4.下载协程：从管道里读取链接并下载

type Task struct {
	url   string // 整页链接
	count int    // 图片数量
}

var (
	chanImageUrls chan string // 存放图片链接的数据管道
	chanTask      chan Task   // 用于监控协程
	waitGroup     sync.WaitGroup
	reImg         = `https?:\/\/[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
)

func main() {

	// 1.初始化管道
	chanImageUrls = make(chan string, 1000000)
	chanTask = make(chan Task, 26)
	// 2.爬虫协程
	for i := 1; i < 27; i++ {
		waitGroup.Add(1)
		go getImgUrls("https://www.bizhizu.cn/shouji/tag-%E5%8F%AF%E7%88%B1/" + strconv.Itoa(i) + ".html")
	}
	// 3.任务统计协程，统计10个任务是否都完成，完成则关闭管道
	waitGroup.Add(1)
	go CheckOK()
	// 4.下载协程：从管道中读取链接并下载
	for i := 0; i < 5; i++ {
		waitGroup.Add(1)
		go DownloadImg()
	}
	waitGroup.Wait()
}

// 下载图片
func DownloadImg() {
	for url := range chanImageUrls {
		filename := GetFilenameFromUrl(url)
		ok := DownloadFile(url, filename)
		if ok {
			fmt.Printf("%s 下载成功\n", filename)
		} else {
			fmt.Printf("%s 下载失败\n", filename)
		}
	}
	waitGroup.Done()
}

// 截取url名字
func GetFilenameFromUrl(url string) (filename string) {
	// 返回最后一个/的位置
	lastIndex := strings.LastIndex(url, "/")
	// 切出来
	filename = url[lastIndex+1:]
	// 时间戳解决重名
	timePrefix := strconv.Itoa(int(time.Now().UnixNano()))
	filename = timePrefix + "_" + filename
	return
}

// 下载图片，传入的是图片叫什么
func DownloadFile(url string, filename string) (ok bool) {
	resp, err := http.Get(url)
	HandleError(err, "http.Get")
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	HandleError(err, "io.ReadAll")
	filename = "./../../assets/crawler/" + filename
	// 写出数据
	err = os.WriteFile(filename, bytes, 0666)
	if err != nil {
		return false
	} else {
		return true
	}
}

// 任务统计协程
func CheckOK() {
	var count int
	for {
		task := <-chanTask
		fmt.Printf("%s 完成了爬取任务，共找到%d条结果\n", task.url, task.count)
		count++
		if count == 26 {
			close(chanImageUrls)
			break
		}
	}
	waitGroup.Done()
}

// 爬图片链接到管道
// url是传的整页链接
func getImgUrls(url string) {
	urls := getImgs(url)
	// 遍历切片里所有链接，存入数据管道
	for _, url := range urls {
		chanImageUrls <- url
	}
	// 标识当前协程完成
	// 每完成一个任务，写一条数据
	// 用于监控协程知道已经完成了几个任务
	chanTask <- Task{url, len(urls)}
	waitGroup.Done()
}

// 获取当前页图片链接
func getImgs(url string) (urls []string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reImg)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		url := result[0]
		urls = append(urls, url)
	}
	return urls
}

// 抽取根据url获取内容
func GetPageStr(url string) (pageStr string) {
	resp, err := http.Get(url)
	HandleError(err, "http.Get")
	defer resp.Body.Close()
	// 读取页面内容
	pageBytes, err := io.ReadAll(resp.Body)
	HandleError(err, "io.ReadAll")
	// 字节转字符串
	pageStr = string(pageBytes)
	return pageStr
}

func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}
