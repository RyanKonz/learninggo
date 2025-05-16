package main

import (
	"errors"
	"fmt"
	"strings"
)

var ErrInvalidID = errors.New("invalid id")

type EmptyFieldError struct {
	EmployeeName string
	Message      string
}

func (ce EmptyFieldError) Error() string {
	return ce.Message
}

func main() {
	err := ReturnEmployeeError()
	var message string = ""
	if err != nil {
		switch err := err.(type) {
		case interface{ Unwrap() []error }:
			allErrors := err.Unwrap()
			var messages []string
			for _, e := range allErrors {
				messages = append(messages, processError(e))
			}
			message = message + " allErrors: " + strings.Join(messages, ", ")
		default:
			message = message + " error: " + processError(err)
		}
	}
	fmt.Println(message)
}

func processError(err error) string {
	var fieldErr EmptyFieldError
	if errors.Is(err, ErrInvalidID) {
		return fmt.Sprintf("invalid ID: %s", err)
	} else if errors.As(err, &fieldErr) {
		return fmt.Sprintf("empty field %s", fieldErr.EmployeeName)
	} else {
		return fmt.Sprintf("%v", err)
	}
}

func ReturnEmployeeError() error {
	var allErrors []error

	allErrors = append(allErrors, EmptyFieldError{
		Message:      "Employee name empty error",
		EmployeeName: "John",
	})

	allErrors = append(allErrors, EmptyFieldError{
		Message:      "Employee name empty error",
		EmployeeName: "Smith",
	})

	switch len(allErrors) {
	case 0:
		return nil
	case 1:
		return allErrors[0]
	default:
		return errors.Join(allErrors...)
	}
}
