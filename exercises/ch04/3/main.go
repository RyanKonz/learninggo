package main

import "fmt"

func main() {
	var total int
	for i := 0; i < 10; i++ {
		total := total + i // total is zero each time here because of shadowing. so the value of i is set the shadowed total and therefore increments up one. If you remove the := and use =, total increases based on previous total.
		fmt.Println(total)
	}
	fmt.Println(total) // this prints 0 because it is the zero value of an int. It is not 9 because the total in the for loop is shadowing and block exists only in the loop.
}
