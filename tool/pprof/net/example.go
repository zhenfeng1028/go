package main

import (
	"io"
	"net/http"
	_ "net/http/pprof"
)

// 导入net/http/pprof包，注意该包利用下划线"_"导入，意味着我们只需要该包运行其init()函数即可，如此该包将自动完成信息采集并保存在内存中。
// 在服务上线时需要将net/http/pprof包移除，其不仅影响服务的性能，更重要的是会造成内存的不断上涨。

func main() {
	h := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc!\n")
	}
	http.HandleFunc("/get", h)
	http.ListenAndServe("localhost:8000", nil)
}

// go build example.go
// ./example

// 可视化
// go tool pprof main http://localhost:8000/debug/pprof/heap
// (pprof) web
