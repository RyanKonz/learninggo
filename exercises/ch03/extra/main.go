package main

import (
	"fmt"
	"slices"
)

func main() {
	//array
	array1 := [3]int{}
	fmt.Println("array1: ", array1)
	fmt.Printf("length of array1: %v\n", len(array1))
	array1[0] = 1
	array1[1] = 2
	array1[2] = 3
	fmt.Println("array1: ", array1)
	fmt.Printf("length of array1: %v\n", len(array1))

	//slice
	sliceFromArray := array1[:]
	fmt.Println("sliceFromArray: ", sliceFromArray)
	fmt.Printf("length of sliceFromArray: %v\n", len(sliceFromArray))
	fmt.Printf("cap of sliceFromArray: %v\n", cap(sliceFromArray))
	sliceFromArray = append(sliceFromArray, 4)
	fmt.Println("sliceFromArray: ", sliceFromArray)
	fmt.Printf("length of sliceFromArray: %v\n", len(sliceFromArray))
	fmt.Printf("cap of sliceFromArray: %v\n", cap(sliceFromArray))

	slice1 := []int{}
	slice1 = append(slice1, 1, 2, 3, 4)
	fmt.Println("slice1: ", slice1)
	fmt.Printf("length of slice1: %v\n", len(slice1))
	fmt.Printf("cap of slice1: %v\n", cap(slice1))

	copyOfSlice1 := make([]int, 4)
	copy(copyOfSlice1, slice1)
	fmt.Println("copyOfSlice1: ", copyOfSlice1)
	fmt.Printf("length of copyOfSlice1: %v\n", len(copyOfSlice1))
	fmt.Printf("cap of copyOfSlice1: %v\n", cap(copyOfSlice1))

	exclude2Slice := slices.DeleteFunc(copyOfSlice1, func(i int) bool {
		return i == 2
	})
	fmt.Println("exclude2Slice: ", exclude2Slice)
	fmt.Printf("length of exclude2Slice: %v\n", len(exclude2Slice))
	fmt.Printf("cap of exclude2Slice: %v\n", cap(exclude2Slice))

	clear(copyOfSlice1)
	fmt.Println("copyOfSlice1: ", copyOfSlice1)
	fmt.Printf("length of copyOfSlice1: %v\n", len(copyOfSlice1))
	fmt.Printf("cap of copyOfSlice1: %v\n", cap(copyOfSlice1))

	//map
	map1 := map[string][]string{
		"One": []string{"A"},
		"Two": []string{"B"},
	}
	fmt.Println("map1: ", map1)

	map1StringSlice := map1["One"]
	fmt.Println("map1StringSlice: ", map1StringSlice)

	nope := map1["Three"] // no panic
	fmt.Println("nope: ", nope)

	_, ok := map1["Three"] // ok idiom
	if !ok {
		fmt.Println("Three not present!")
	}

	makeMap := make(map[string]string, 5)
	fmt.Println("makeMap: ", makeMap)
	fmt.Printf("length of makeMap: %v\n", len(makeMap))

	makeMap["One"] = "1"
	fmt.Println("makeMap: ", makeMap)
	fmt.Printf("length of makeMap: %v\n", len(makeMap))

	makeMap["Two"] = "2"
	makeMap["Three"] = "3"
	makeMap["Four"] = "4"
	makeMap["Five"] = "5"
	makeMap["Six"] = "6" // no panic
	fmt.Println("makeMap: ", makeMap)
	fmt.Printf("length of makeMap: %v\n", len(makeMap))

	makeMap["Four"] = "Four"
	fmt.Println("makeMap: ", makeMap)
	fmt.Printf("length of makeMap: %v\n", len(makeMap))

	delete(makeMap, "Four")
	fmt.Println("makeMap: ", makeMap)
	fmt.Printf("length of makeMap: %v\n", len(makeMap))

	clear(makeMap)
	fmt.Println("makeMap: ", makeMap)
	fmt.Printf("length of makeMap: %v\n", len(makeMap))

	//map as set
	intSet := map[int]bool{}
	vals := 1
	intSet[vals] = true

	if intSet[vals] {
		fmt.Println("vals exists")
	}

	// anon structs
	person := struct {
		firstName string
		lastName  string
	}{}
	fmt.Println("person: ", person)

	person.firstName = "i am"
	person.lastName = "groot"
	fmt.Println("person: ", person)

	var employee struct {
		firstName string
		lastName  string
	}
	employee.firstName = "i am"
	fmt.Println("employee: ", employee)
}
