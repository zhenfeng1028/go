package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {

	var nameFile string
	flag.StringVar(&nameFile, "f", "name.json", "name file path")
	flag.Parse()

	fmt.Println(nameFile)

	nameData, err := ioutil.ReadFile(nameFile)
	if err != nil {
		fmt.Printf("config file read err: %s ", err)
	}
	fmt.Printf("config load:\n %s", string(nameData))
}
