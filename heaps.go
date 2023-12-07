package main

import (
	"fmt"
	"log"
)

// MaxHeap struct has a slice that holds the array
type MaxHeap struct {
	array []int
}

// Insert adds an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// Extract return the largest key, and removes it from the heap
func (h *MaxHeap) Extract() int {
	// Return the extracted item
	extracted := h.array[0]

	length := len(h.array) - 1

	if len(h.array) == 0 {
		log.Fatal("Cannot extract! Bacause is a invalid heap with 0 items inside")
		return -1
	}

	// Get the last key index and put into the first index (Root)
	h.array[0] = h.array[length]
	// Now here is a slice method
	h.array = h.array[:length]

	h.maxHeapifyDown(0)

	return extracted
}

// Swap the values the current index is larger or biger than parent
func (h *MaxHeap) maxHeapifyUp(index int) {
	// MaxHeapify will heapify from brottom to top
	for h.array[parent(index)] < h.array[index] {
		// This part is swaping until it find its place
		h.swap(parent(index), index)
		// Update value
		index = parent(index)
	}
}

// MaxHeapifyDown will heapify top to bottom
func (h *MaxHeap) maxHeapifyDown(index int) {
	// Declare the variables to compare
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	// Loop while index has at least one child
	for l <= lastIndex {
		if l == lastIndex {
			// When left child is the only child
			childToCompare = l
		} else if h.array[l] > h.array[r] {
			// When left child is larger
			childToCompare = l
		} else {
			// When the right child is larger
			childToCompare = r
		}

		// Compare array value of current index to larger child and swap if smaller
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			// If the current index is actually larger than larger child that means that he found the correct place
			return
		}
	}
}

// Get the parent index value
// i = parent
func parent(i int) int {
	return (i - 1) / 2
}

// Get the left child index value
// i = parent
func left(i int) int {
	return 2*i + 1
}

// Get the right child index value
// i = parent
func right(i int) int {
	return 2*i + 2
}

// Swap keys in the array
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {
	// Creating a Heap (& => Address variable)
	m := &MaxHeap{}
	fmt.Println(m)

	// Add some keys
	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}
