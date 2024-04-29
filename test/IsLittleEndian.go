package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println(IsLittleEndian())
}

func IsLittleEndian() bool {
	var n int32 = 0x01020304

	// 下面是为了将int32类型的指针转换成byte类型的指针
	u := unsafe.Pointer(&n)
	pb := (*byte)(u)

	// 取得pb位置对应的值
	b := *pb

	// 由于b是byte类型，最多保存8位，那么只能取得开始的8位
	// 小端: 04 (03 02 01)
	// 大端: 01 (02 03 04)
	return (b == 0x04)
}
