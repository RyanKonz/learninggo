package main

import (
	"fmt"
	"math/rand"
)

func main() {
	slice := make([]int, 0, 100)
	for range 100 {
		slice = append(slice, rand.Intn(100))
	}
	fmt.Println(slice)

	for _, v := range slice {
		//fmt.Println(v)
		switch {
		case v%6 == 0:
			fmt.Println("Six!")
		case v%2 == 0:
			fmt.Println("Two!")
		case v%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Never mind")
		}

		if v%2 == 0 && v%3 == 0 {
			fmt.Println("Six!")
		} else if v%2 == 0 {
			fmt.Println("Two!")
		} else if v%3 == 0 {
			fmt.Println("Three!")
		} else {
			fmt.Println("Never mind")
		}
	}
}
