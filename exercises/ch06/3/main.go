package main

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

const countOfPersons = 10_000_000

func main() {
	p := make([]Person, 0, countOfPersons)
	//var p []Person
	for range countOfPersons {
		p = append(p, Person{"john", "smith", 40})
	}
}
