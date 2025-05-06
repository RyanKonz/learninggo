package main

import "fmt"

type ValidTypes interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func doubler[T ValidTypes](input T) T {
	return input * 2
}

func main() {
	fmt.Println(doubler(5))
	fmt.Println(doubler(int64(10)))
	fmt.Println(doubler(float32(10.10)))
}
