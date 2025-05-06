package main

import "fmt"

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

type List[T comparable] struct {
	Head *Node[T]
	Tail *Node[T]
}

func (l *List[T]) Add(t T) {
	n := &Node[T]{
		Value: t,
	}

	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}

	l.Tail.Next = n
	l.Tail = l.Tail.Next
}

func (l *List[T]) Insert(t T, position int) {
	n := &Node[T]{
		Value: t,
	}

	// if empty, insert
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}

	// if position is first, make new node first
	if position <= 0 {
		n.Next = l.Head
		l.Head = n
		return
	}

	// curNode will traverse list
	curNode := l.Head
	for i := 1; i < position; i++ {
		// if position is past the length, add it to the end and return
		if curNode.Next == nil {
			curNode.Next = n
			l.Tail = curNode.Next
			return
		}
		curNode = curNode.Next
	}

	// insert
	n.Next = curNode.Next
	curNode.Next = n

	// if curNode is the tail, make the new node the tail
	if l.Tail == curNode {
		l.Tail = n
	}

	return
}

func (l *List[T]) Index(t T) int {
	i := 0
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		if curNode.Value == t {
			return i
		}
		i++
	}
	return -1
}

func main() {
	l := &List[int]{}
	l.Add(5)
	l.Add(10)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(20))

	l.Insert(100, 0)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(20))
	fmt.Println(l.Index(100))

	l.Insert(200, 1)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(200))
	fmt.Println(l.Index(20))
	fmt.Println(l.Index(100))

	fmt.Println("printing all")
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		fmt.Println(curNode.Value)
	}

	l.Insert(300, 10)
	fmt.Println("printing all")
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		fmt.Println(curNode.Value)
	}

	l.Add(400)
	fmt.Println("printing all")
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		fmt.Println(curNode.Value)
	}

	l.Insert(500, 6)
	fmt.Println("printing all")
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		fmt.Println(curNode.Value)
	}
}
