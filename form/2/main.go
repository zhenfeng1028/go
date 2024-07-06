package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析url传递的参数，对于POST则解析响应包的主体（request body）
	// 注意：如果没有调用ParseForm方法，下面无法获取表单的数据
	fmt.Println("form:", r.Form) // 这些信息是输出到服务器端的打印信息
	fmt.Println("path:", r.URL.Path)
	fmt.Println("scheme:", r.URL.Scheme)
	fmt.Fprintf(w, "Hello astaxie!") // 这个写入到w的是输出到客户端的
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // 获取请求的方法
	if r.Method == "GET" {
		t, err := template.ParseFiles("login.gtpl")
		if err != nil {
			log.Fatalf("parse file err: %v\n", err)
		}
		log.Println(t.Execute(w, nil))
	} else {
		// 请求的是登录数据，那么执行登录的逻辑判断
		r.ParseForm()
		bs, err := json.Marshal(r.Form)
		if err != nil {
			log.Fatalf("marshal err: %v\n", err)
		}
		fmt.Println("r.Form:", string(bs))
		fmt.Println("username:", r.FormValue("username"))
		fmt.Println("password:", r.FormValue("password"))
	}
}

func main() {
	http.HandleFunc("/", sayhelloName)       // 设置访问的路由
	http.HandleFunc("/login", login)         // 设置访问的路由
	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
