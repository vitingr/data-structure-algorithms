package main

// The main diference between a stack and a queue is the way that the data is removed

// -=-=-=-=-=-=-=- STACKS -=-=-=-=-=-=-=-=-=-
// Stacks have a Last in First Out data structure (LIFO)
// That the last one that you put in, would be the one that you take out first
// Like a book stack. We use PUSH => to insert items and POP => to remove items

import (
	"fmt"
)

// Stack represents a stack that hold a slice
type Stack struct {
	items []int
}

// Push items from Stack
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop items from Stack (will remove a value at the end and return the removed value)
func (s *Stack) Pop() int {
	l := len(s.items)-1

	toRemove := s.items[l]
	s.items = s.items[:l]

	return toRemove
}

func main() {
	myStack := Stack{}

	fmt.Println(myStack)

	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)

	fmt.Println(myStack)

	myStack.Pop()

	fmt.Println(myStack)
}
