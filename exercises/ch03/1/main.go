package main

import "fmt"

func main() {
	greetings := []string{
		"Hello",
		"Hola",
		"नमस्कार",
		"こんにちは",
		"Привіт",
	}

	first := greetings[:2]
	second := greetings[1:4]
	third := greetings[3:]

	fmt.Println(greetings)
	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)

	//fmt.Printf("%v : %v", greetings, len(greetings))
}
