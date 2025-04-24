package main

import (
	"fmt"
	"math/rand"
)

func main() {
	slice := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		slice = append(slice, rand.Intn(100))
	}
	fmt.Println(slice)

	slice2 := make([]int, 0, 100)
	for range 100 {
		slice2 = append(slice2, rand.Intn(100))
	}
	fmt.Println(slice2)
}
