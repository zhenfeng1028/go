package main

import (
	"strings"
	"sync"
)

var (
	populateMapsOnce sync.Once
	entity           map[string]rune
)

func populateMaps() {
	entity = map[string]rune{
		"AElig;":  '\U000000C6',
		"AMP;":    '\U00000026',
		"Aacute;": '\U000000C1',
		"Abreve;": '\U00000102',
		"Acirc;":  '\U000000C2',
		// 省略 2000 项
	}
}

func UnescapeString(s string) string {
	populateMapsOnce.Do(populateMaps)
	i := strings.IndexByte(s, '&')

	if i < 0 {
		return s
	}
	// 省略后续的实现
	return "string after unescaped"
}

// 字典 entity 包含 2005 个键值对，若使用 init 在包加载时初始化，若不被使用，将会浪费大量内存。
// html.UnescapeString(s) 函数是线程安全的，可能会被用户程序在并发场景下调用，因此对 entity 的初始化需要加锁，使用 sync.Once 能保证这一点。
