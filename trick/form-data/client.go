package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	// 服务器地址
	url := "http://localhost:8080/upload"

	// 创建一个缓冲区用于存储 multipart/form-data 数据
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// 添加文本字段
	err := writer.WriteField("username", "testuser")
	if err != nil {
		fmt.Println("Error writing field:", err)
		return
	}

	// 添加文件字段
	file, err := os.Open("output.txt") // 替换为你要上传的文件路径
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	part, err := writer.CreateFormFile("file", "output.txt") // 字段名和文件名
	if err != nil {
		fmt.Println("Error creating form file:", err)
		return
	}

	_, err = io.Copy(part, file) // 将文件内容写入 multipart 表单
	if err != nil {
		fmt.Println("Error copying file content:", err)
		return
	}

	// 关闭 multipart writer
	writer.Close()

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", url, &requestBody)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// 设置 Content-Type 为 multipart/form-data
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", string(responseBody))
}
