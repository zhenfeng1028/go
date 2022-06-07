package main

import (
	"fmt"
	"regexp"
)

func main() {
	p := regexp.MustCompile(`a.`)
	fmt.Println(p.Find([]byte("ababab")))
	fmt.Println(p.FindString("ababab"))
	fmt.Println(p.FindAllString("ababab", -1))
	fmt.Println(p.FindAllStringIndex("ababab", -1))

	q, _ := regexp.Compile(`^a(.*)b$`)
	fmt.Println(q.FindAllSubmatch([]byte("ababab"), -1))
	fmt.Println(q.FindAllStringSubmatch("ababab", -1))
	fmt.Println(q.FindAllStringSubmatchIndex("ababab", -1))

	r := regexp.MustCompile(`(?m)(key\d+):\s+(value\d+)`)
	content := []byte(`
        # comment line
        key1: value1
        key2: value2
        key3: value3
    `)
	fmt.Println(string(r.Find(content)))
	for _, matched := range r.FindAll(content, -1) {
		fmt.Println(string(matched))
	}
	for _, mutiMatched := range r.FindAllSubmatch(content, -1) {
		for _, matched := range mutiMatched {
			fmt.Println(string(matched))
		}
	}
}
