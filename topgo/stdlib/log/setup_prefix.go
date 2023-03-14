// log标准库中还提供了关于日志信息前缀的两个方法：
//     func Prefix() string
//     func SetPrefix(prefix string)
// 其中Prefix函数用来查看标准logger的输出前缀，SetPrefix函数用来设置输出前缀。

package main

import "log"

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate)
	log.Println("这是一条很普通的日志。")
	log.SetPrefix("[Info]")
	log.Println("这是一条很普通的日志。")
}
