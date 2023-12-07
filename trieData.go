package main

import (
	"fmt"
)

// A Trie is a tree structure and each node represent a word or a part of a word
// Path conected by the root can store words too
// An exemple of TrieData is the google searchbar, that tries to guess what you are trying to search
// Each node will have 26 children nodes

const AlphabetSize = 26

// Node Structure represents each node in the trie
// Each index of that array will hold a pointer to the chlild
type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

// Trie Structure represent a trie and has a pointer to the root node
type Trie struct {
	root *Node
}

// InitTrie will create a new Trie
func InitTrie() *Trie {
	// Create a address variable
	result := &Trie{root: &Node{}}
	return result
}

// Insert will take in a word and add it to the trie
func (t *Trie) Insert(w string) {
	wordLength := len(w)
	currentNode := t.root

	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}

	currentNode.isEnd = true
}

// Search will take in a word and return true is that word is included in the trie
func (t *Trie) Search(w string) bool {
	wordLength := len(w)
	currentNode := t.root

	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}

	if currentNode.isEnd == true {
		return true
	}

	return false
}

func main() {
	myTrie := InitTrie()

	toAdd := []string{
		"rayani",
		"melissa",
		"daniel",
		"lais",
	}

	for _, value := range toAdd {
		myTrie.Insert(value)
	}

	fmt.Println(myTrie.Search("rayani"))
}
