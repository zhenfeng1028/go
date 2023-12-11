package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
)

func main() {
	src := "hello world"
	compressed := DoZlibCompress([]byte(src))
	fmt.Println(compressed)

	decompressed := DoZlibUnCompress(compressed)
	fmt.Println(string(decompressed))
}

// zlib压缩
func DoZlibCompress(src []byte) []byte {
	var in bytes.Buffer
	w, _ := zlib.NewWriterLevel(&in, zlib.BestCompression)
	w.Write(src)
	w.Close()
	return in.Bytes()
}

// zlib解压缩
func DoZlibUnCompress(compressed []byte) []byte {
	b := bytes.NewReader(compressed)
	var out bytes.Buffer
	r, _ := zlib.NewReader(b)
	io.Copy(&out, r)
	return out.Bytes()
}
