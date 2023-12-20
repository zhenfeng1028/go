package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(krand(4, 0))
}

// 生成随机验证码
// size 验证码位数
// kind 验证码类型 0 纯数字 1 小写字母 2 大写字母 other 数字、大小写字母
func krand(size int, kind int) string {
	ikind, kinds, result := kind, [][]int{{10, 48}, {26, 97}, {26, 65}}, make([]byte, size)
	isAll := kind > 2 || kind < 0
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := 0; i < size; i++ {
		if isAll { // random ikind
			ikind = r.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}
