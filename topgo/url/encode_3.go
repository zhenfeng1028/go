package main

import (
	"fmt"
	"net/url"
)

func main() {
	// Let's start with a base url
	baseUrl, err := url.Parse("http://www.mywebsite.com")
	if err != nil {
		fmt.Println("Malformed URL: ", err.Error())
		return
	}
	// Add a Path Segment (Path segment is automatically escaped)
	baseUrl.Path += "path with?reserved characters"
	// Prepare Query Parameters
	params := url.Values{}
	params.Add("q", "Hello World")
	params.Add("u", "@rajeev")
	// Add Query Parameters to the URL
	baseUrl.RawQuery = params.Encode() // Escape Query Parameters
	fmt.Printf("Encoded URL is %q\n", baseUrl.String())
}
