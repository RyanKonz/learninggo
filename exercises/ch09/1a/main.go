package main

import "fmt"

type CustomError struct {
	Message string
}

func (ce CustomError) Error() string {
	return ce.Message
}

func main() {
	err := ReturnError()
	if err != nil {
		fmt.Printf("%v", err)
	}
}

func ReturnError() error {
	return CustomError{
		Message: "this is a custom error",
	}
}
