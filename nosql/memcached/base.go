package main

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

var (
	server = "127.0.0.1:11211"
)

func main() {
	// create a handle
	mc := memcache.New(server)
	if mc == nil {
		fmt.Println("memcache New failed")
	}

	// set key-value
	err := mc.Set(&memcache.Item{Key: "foo", Value: []byte("my value")})
	if err != nil {
		fmt.Println("Set failed:", err.Error())
	}

	// get key's value
	it, err := mc.Get("foo")
	if err != nil {
		fmt.Println("Get failed:", err.Error())
	} else {
		fmt.Println("value is", string(it.Value))
	}

	// add a new key-value
	err = mc.Add(&memcache.Item{Key: "foo", Value: []byte("bluegogo")})
	if err != nil {
		fmt.Println("Add failed:", err.Error())
	}
	it, err = mc.Get("foo")
	if err != nil {
		fmt.Println("Get failed:", err.Error())
	} else {
		fmt.Println("value is", string(it.Value))
	}

	// replace a key's value
	err = mc.Replace(&memcache.Item{Key: "foo", Value: []byte("mobike")})
	if err != nil {
		fmt.Println("Replace failed:", err.Error())
	}
	it, err = mc.Get("foo")
	if err != nil {
		fmt.Println("Get failed:", err.Error())
	} else {
		fmt.Println("Replace value is", string(it.Value))
	}

	// delete an exist key
	err = mc.Delete("foo")
	if err != nil {
		fmt.Println("Delete failed:", err.Error())
	}

	// incrby
	err = mc.Set(&memcache.Item{Key: "aaa", Value: []byte("1")})
	if err != nil {
		fmt.Println("Set failed:", err.Error())
	}
	it, err = mc.Get("aaa")
	if err != nil {
		fmt.Println("Get failed:", err.Error())
	} else {
		fmt.Println("src value is", string(it.Value))
	}
	value, err := mc.Increment("aaa", 7)
	if err != nil {
		fmt.Println("Increment failed:", err.Error())
	} else {
		fmt.Println("after increment the value is", value)
	}

	// decrby
	value, err = mc.Decrement("aaa", 4)
	if err != nil {
		fmt.Println("Decrement failed:", err.Error())
	} else {
		fmt.Println("after decrement the value is", value)
	}
}
