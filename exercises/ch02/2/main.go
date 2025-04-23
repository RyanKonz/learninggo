package main

import "fmt"

func main() {
	const value = 1
	var i int = value
	var f float32 = value
	fmt.Println(i)
	fmt.Println(f)

	a := value
	fmt.Println(a)
}
