package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const str = `&<>`

func main() {
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	if err := jsonEncoder.Encode(str); err != nil {
		return
	}
	fmt.Println(bf.String())
}
