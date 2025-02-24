package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	unicodeStr := "\\u6587\\u4ef6\\u7ffb\\u8bd1"
	chineseStr, err := unicodeToChinese(unicodeStr)
	if err != nil {
		fmt.Println("转换失败:", err)
		return
	}
	fmt.Println("转换结果:", chineseStr)

	unicodeStr = chineseToUnicode("文件翻译")
	fmt.Println("转换结果:", unicodeStr)
}

func unicodeToChinese(unicodeStr string) (string, error) {
	parts := strings.Split(unicodeStr, "\\u")
	var result strings.Builder

	for _, part := range parts {
		if len(part) == 0 {
			continue
		}
		code, err := strconv.ParseInt(part, 16, 32)
		if err != nil {
			return "", fmt.Errorf("无效的 Unicode 编码: %s", part)
		}
		result.WriteRune(rune(code))
	}

	return result.String(), nil
}

func chineseToUnicode(chineseStr string) string {
	var result string
	for _, runeValue := range chineseStr {
		result += fmt.Sprintf("\\u%04x", runeValue)
	}
	return result
}
