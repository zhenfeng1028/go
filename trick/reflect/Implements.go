package main

import (
	"fmt"
	"reflect"
)

// Define an interface
type Writer interface {
	Write([]byte) (int, error)
}

// Define a type that implements the Writer interface
type FileWriter struct{}

func (f FileWriter) Write(data []byte) (int, error) {
	fmt.Println("Writing data:", string(data))
	return len(data), nil
}

// Define a type that does NOT implement the Writer interface
type Reader struct{}

func (r Reader) Read(data []byte) (int, error) {
	fmt.Println("Reading data:", string(data))
	return len(data), nil
}

func main() {
	// Create instances of the types
	fileWriter := FileWriter{}
	reader := Reader{}

	// Get the reflect.Type of the instances
	fileWriterType := reflect.TypeOf(fileWriter)
	readerType := reflect.TypeOf(reader)

	// Get the reflect.Type of the Writer interface
	writerInterface := reflect.TypeOf((*Writer)(nil)).Elem()

	// Check if FileWriter implements the Writer interface
	if fileWriterType.Implements(writerInterface) {
		fmt.Println("FileWriter implements the Writer interface")
	} else {
		fmt.Println("FileWriter does NOT implement the Writer interface")
	}

	// Check if Reader implements the Writer interface
	if readerType.Implements(writerInterface) {
		fmt.Println("Reader implements the Writer interface")
	} else {
		fmt.Println("Reader does NOT implement the Writer interface")
	}
}
