package main

// The main diference between a stack and a queue is the way that the data is removed

// -=-=-=-=-=-=-=- QUEUES -=-=-=-=-=-=-=-=-=-
// Stacks have a First in First Out data struct (FIFO)
// Like a supermarket cashier, that there is a queue, and is linear
// The last item goes to the last index, and the first item would be the first
// Like a book stack. We use ENQUEUE => to insert items and DEQUEUE => to remove items

import (
	"fmt"
)

type Queue struct {
	items []int
}

// Enqueue method (will add a value into the current Queue)
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Dequeue method (will remove a value at the front and return the removed value)
func (q *Queue) Dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]

	return toRemove
}

func main() {
	myQueue := Queue{}
	fmt.Println(myQueue)

	myQueue.Enqueue(100)
	myQueue.Enqueue(200)
	myQueue.Enqueue(300)

	fmt.Println(myQueue)

	myQueue.Dequeue()

	fmt.Println(myQueue)
}