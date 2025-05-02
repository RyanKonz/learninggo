package main

type Team struct {
	name    string
	players []string
}

type League struct {
	Teams map[string]Team
	Wins  map[string]int
}

func main() {
}
