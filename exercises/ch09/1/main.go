package main

import (
	"errors"
	"fmt"
)

var ErrInvalidId = errors.New("invalid id error")

func main() {
	err := ReturnError(0)
	if err != nil {
		if errors.Is(ErrInvalidId, err) {
			fmt.Printf("%v\n", err)
		} else {
			fmt.Println(err)
		}
	}
}

func ReturnError(i int) error {
	if i == 0 {
		return ErrInvalidId
	}
	return errors.New("not invalid id")
}
