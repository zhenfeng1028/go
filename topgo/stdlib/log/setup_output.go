// SetOutput函数用来设置标准logger的输出目的地，默认是标准错误输出。
// func SetOutput(w io.Writer)
// 例如，下面的代码会把日志输出到同目录下的xx.log文件中。

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	Init()
	// logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	// if err != nil {
	// 	fmt.Println("open log file failed, err: ", err)
	// 	return
	// }
	// log.SetOutput(logFile)
	// log.SetFlags(log.Lshortfile | log.Ldate)
	log.Println("这是一条很普通的日志。")
	log.SetPrefix("[Info]")
	log.Println("这是一条很普通的日志。")
}

// 如果你要使用标准的logger，我们通常会把上面的配置操作写到init函数中。
func Init() {
	logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err: ", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
}
