package main

import (
	"fmt"
	"log"
)

// -=-=-=-=-=-=-=- HASH TABLES -=-=-=-=-=-=-=-=-=-
// Imagine a big array, to search a value into the entire array will late a long time
// In the hashtables is faster because you search directly by index that is generated uby a hash function
// Take care with collision handling, 2 values for the same variable index
// Open Address, if there is a index into a hash variable, the next item will be placed into the next index
// You can put a linked list inside a hash variable (separate chaining)
// Buckets are similiar with the childs of the hashtables variables

// Array size is the size of the hash table array
const ArraySize = 7

// Hashtable Structure will hold an array
type HashTable struct {
	array [ArraySize]*bucket
}

// Bucket Structure is a linked list in each slot of the hash table
type bucket struct {
	head *bucketNode
}

// BucketNode structure
type bucketNode struct {
	key  string
	next *bucketNode
}

// Insert will take in a key and add it to the hash table array
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// Search will take in a key and return true if that key is stored in the hash table
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

// Delete will take in a key and delete it from the hash table
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

// Insert bucket will take in a key, create a node with the key and insert the node in the bucket
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		log.Fatal("Cannot insert a value that already exists into this hash table")
	}
}

// Search bucket will take in a key and return true if the bucket has that key
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// Delete bucket will take in a key and delete the node from the bucket
func (b *bucket) delete(k string) {

	if b.head.key == k {
		b.head = b.head.next
		return
	}

	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			// Delete
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}
}

// Hash
func hash(key string) int {
	sum := 0
	for _, value := range key {
		sum += int(value)
	}
	return sum % ArraySize
}

// Init will create a bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}

	// To go to all slots of the hash table array
	for i := range result.array {
		// Inside each array created a bucket
		result.array[i] = &bucket{}
	}

	return result
}

func main() {
	// Declare and init the hashtable variable
	hashTable := Init()

	// Array of items to be added to hash
	list := []string{
		"GETULIO",
		"PATRICIO",
		"ALFREDO",
		"SEBASTIÃO",
		"HORÁCIO",
		"NATALINO",
		"ALFINAS",
	}
	
	// Method to add all items from list
	for _, value := range list {
		hashTable.Insert(value)
	}

	hashTable.Delete("ALFINAS")
	fmt.Println("ALFREDO", hashTable.Search("ALFREDO"))
}
