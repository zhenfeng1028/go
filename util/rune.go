package main

import "fmt"

func main() {
	s := "abc你好"
	r := "123你好"
	fmt.Println("len(s)=", len([]byte(s)), "len(r)=", len([]rune(r))) // len(s)= 9 len(r)= 5
	fmt.Printf("%x %x %x %x %x %x\n", r[0], r[1], r[2], r[3], r[4], r[5])
	for k, v := range r {
		fmt.Println("k=", k, "v=", v)
	}

	for k, v := range []rune(r) {
		fmt.Printf("k2=%d v2=%d %x\n", k, v, v)
	}
}

// golang中string底层是通过byte数组实现的。中文字符在unicode下占2个字节，在utf-8编码下占3个字节，golang默认编码是utf-8。

// 如果想得到字符串真正的长度(一个中文算一位)，需要将字符串转换为rune，再求长度。

// 但是用range遍历包含中文的字符串时，会发现第一个for中k的值为0,1,2,3,6；第二个for中k的值为0,1,2,3,4；
// 说明字符串中如果包含中文，range的时候可以识别出来，一个中文占用3个byte，索引就自动加三；而rune索引则只会加一。
