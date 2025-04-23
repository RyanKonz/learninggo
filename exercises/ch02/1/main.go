package main

import "fmt"

func main() {
	var i int = 20
	var f float32 = float32(i)
	fmt.Println(i)
	fmt.Println(f)

	a := 20
	b := float32(a)
	fmt.Println(a)
	fmt.Println(b)

	var c = 20
	var d = float32(c)
	fmt.Println(c)
	fmt.Println(d)
}
