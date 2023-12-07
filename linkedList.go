package main

import (
	"fmt"
	"log"
)

// Describe the node (cadeia)
type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

// Increase item to the linkedList method
// We declare the types of methods outside the function in POO programming languages
func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedList) printListData() {
	toPrint := l.head

	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
}

// Put the variable into the head to delete
func (l *linkedList) deleteWithValue(value int) {
	// Check if the list is empty
	if l.length == 0 {
		log.Fatal("Cannot delete a value of a empty list, sorry")
		return
	}

	// If the parent value is equal to Value Param
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	// We are not going to compare the date of the node to delete
	// We want to compare the value of the next
	for previousToDelete.next.data != value {
		// Realize a for until find the value to be deleted
		if previousToDelete.next.next == nil {
			log.Fatal("Is not possible to remove a value that doesn't exist")
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main() {
	myList := linkedList{}

	// Put the node values into the linked List
	node1 := &node{data: 48}
	myList.prepend(node1)

	node2 := &node{data: 18}
	myList.prepend(node2)

	node3 := &node{data: 16}
	myList.prepend(node3)

	node4 := &node{data: 11}
	myList.prepend(node4)

	node5 := &node{data: 7}
	myList.prepend(node5)

	node6 := &node{data: 2}
	myList.prepend(node6)

	// Will return the current address of the list into computer memorie
	myList.printListData()
	myList.deleteWithValue(16)
	myList.printListData()
}
