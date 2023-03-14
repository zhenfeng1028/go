// 终端其实是一个文件，相关实例如下：
// os.Stdin：标准输入的文件实例，类型为*File
// os.Stdout：标准输出的文件实例，类型为*File
// os.Stderr：标准错误输出的文件实例，类型为*File
// 以文件的方式操作终端:

package main

import "os"

func main() {
	var buf [16]byte
	os.Stdin.Read(buf[:])
	os.Stdin.WriteString(string(buf[:]))
}
