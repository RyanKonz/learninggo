package main

import "fmt"

func main() {
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	first := Employee{
		"i am",
		"groot",
		1,
	}
	second := Employee{
		id:        2,
		firstName: "i am",
		lastName:  "groot",
	}
	var third Employee
	third.firstName = "i am"
	third.lastName = "groot"
	third.id = 3

	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)

	fourth := Employee{id: 4}
	fmt.Println(fourth)

	fifth := Employee{}
	fmt.Println(fifth)
}
