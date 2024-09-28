package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	// NonceSymbols 随机字符串可用字符集
	NonceSymbols = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// NonceLength 随机字符串的长度
	NonceLength = 32
)

func main() {
	fmt.Println(GetRandomString(32))
	fmt.Println(GenerateNonce())
}

// 获取随机字符串
// length 字符串长度
func GetRandomString(length int) string {
	var (
		result []byte
		b      []byte
		r      *rand.Rand
	)
	b = []byte(NonceSymbols)
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, b[r.Intn(len(b))])
	}
	return string(result)
}

func GenerateNonce() string {
	bytes := make([]byte, NonceLength)
	rand.Read(bytes)
	symbolsByteLength := byte(len(NonceSymbols))
	for i, b := range bytes {
		bytes[i] = NonceSymbols[b%symbolsByteLength]
	}
	return string(bytes)
}
