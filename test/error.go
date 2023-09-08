package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	Msg string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Message: %s", e.Msg)
}

func msgError(msg string) error {
	return &MyError{Msg: msg}
}

func WrapError(msg string) error {
	err := msgError(msg)
	return fmt.Errorf("(Wrapping) %w", err)
}

func main() {
	err1 := msgError("Error")
	fmt.Println("[Normal Error]", err1)

	err2 := WrapError("Test Error Message")
	fmt.Println("[Wrapping Error]", err2)

	var myErr *MyError
	if errors.As(err2, &myErr) {
		fmt.Printf("[Failed] %s\n", myErr)
	}
}
