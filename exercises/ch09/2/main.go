package main

import (
	"errors"
	"fmt"
)

type EmptyFieldError struct {
	EmployeeName string
	Message      string
}

func (ce EmptyFieldError) Error() string {
	return ce.Message
}

func main() {
	err := ReturnEmployeeError()
	var EmptyFieldErr EmptyFieldError
	if err != nil {
		if errors.As(err, &EmptyFieldErr) {
			fmt.Println(EmptyFieldErr.Message, ": ", EmptyFieldErr.EmployeeName)
		}
	}
}

func ReturnEmployeeError() error {
	return EmptyFieldError{
		Message:      "Employee name empty error",
		EmployeeName: "John",
	}
}
