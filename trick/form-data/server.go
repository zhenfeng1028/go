package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// 解析 multipart/form-data 请求
	err := r.ParseMultipartForm(10 << 20) // 限制上传文件大小为 10MB
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// 获取文本字段
	username := r.FormValue("username")
	fmt.Printf("Username: %s\n", username)

	// 获取文件字段
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to get file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// 创建目标文件
	dst, err := os.Create(fileHeader.Filename)
	if err != nil {
		http.Error(w, "Unable to create file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// 将上传的文件内容复制到目标文件
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
		return
	}

	// 返回成功响应
	fmt.Fprintf(w, "File uploaded successfully: %s\n", fileHeader.Filename)
}

func main() {
	// 注册文件上传处理函数
	http.HandleFunc("/upload", uploadHandler)

	// 启动 HTTP 服务器
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
