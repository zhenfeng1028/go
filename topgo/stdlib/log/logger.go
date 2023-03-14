// log包定义了Logger类型，该类型提供了一些格式化输出的方法。
// 本包也提供了一个预定义的“标准”logger，可以通过调用函数
// Print系列（Print|Printf|Println）、
// Fatal系列（Fatal|Fatalf|Fatalln）、
// Panic系列（Panic|Panicf|Panicln）
// 来使用，比自行创建一个logger对象更容易使用。

package main

import (
	"log"
)

func main() {
	log.Println("这是一条很普通的日志。")
	v := "很普通的"
	log.Printf("这是一条%s日志。\n", v)
	log.Fatalln("这是一条会触发fatal的日志。") // Fatal系列函数会在写入日志信息后调用os.Exit(1)。
	log.Panicln("这是一条会触发panic的日志。") // Panic系列函数会在写入日志信息后panic。
}
