package main

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	MakePerson("john", "smith", 40)
	MakePersonPointer("john", "smith", 40)
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}
