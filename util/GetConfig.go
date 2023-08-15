package main

import (
	"encoding/json"
	"fmt"
)

var str = `
{
	"person": {
		"name": "lzf",
		"age": 28
 	}
}`

var globalConfig = make(map[string]interface{})

func main() {
	err := json.Unmarshal([]byte(str), &globalConfig)
	if err != nil {
		panic(err)
	}

	str, _ := GetString("person", "name")
	fmt.Println("person.name:", str)

	v, err := GetInt("person", "age")
	fmt.Println("person.age:", v, err) // 解析出来是float64类型
}

func GetString(keys ...string) (string, error) {
	result, err := Get(globalConfig, keys...)
	if err != nil {
		return "", err
	}

	str, ok := result.(string)
	if !ok {
		return "", fmt.Errorf("type error")
	}
	return str, nil
}

func GetInt(keys ...string) (int, error) {
	result, err := Get(globalConfig, keys...)
	if err != nil {
		return -1, err
	}

	v, ok := result.(int)
	if !ok {
		return -1, fmt.Errorf("type error")
	}
	return v, nil
}

func Get(config interface{}, keys ...string) (interface{}, error) {
	for _, key := range keys {
		c, err := GetKey(config, key)
		if err != nil {
			return nil, err
		}
		config = c
	}
	return config, nil
}

func GetKey(config interface{}, key string) (interface{}, error) {
	m, ok := config.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("config type error")
	}

	result, ok := m[key]
	if !ok {
		return nil, fmt.Errorf("expect key %s", key)
	}
	return result, nil
}
