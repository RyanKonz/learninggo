package main

import "fmt"

func main() {
	s := []string{"a", "b", "c"}
	fmt.Println("before update: ", s)
	UpdateSlice(s, "d")
	fmt.Println("after update: ", s)
	fmt.Println("before grow: ", s)
	GrowSlice(s, "e")
	fmt.Println("after grow: ", s)
}

func UpdateSlice(strSlice []string, str string) {
	strSlice[len(strSlice)-1] = str
	fmt.Println("update slice: ", strSlice)
}

func GrowSlice(strSlice []string, str string) {
	strSlice = append(strSlice, str)
	fmt.Println("grow slice: ", strSlice)
}
