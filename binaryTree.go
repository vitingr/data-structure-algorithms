package main

import (
	"fmt"
)

// -=-=-=-=-=-=-=- TREE -=-=-=-=-=-=-=-=-=-
// Structure of a Tree
// Level 1 = Root (Is most highest into the tree scale, the top)
// Level 2 = Parent (Values below the root)
// Level 3 = Children (Values blow the parents)
// Level 4 = Leaves (Values below the children)
// Always have 2 connections, and the smallest value is always into the left

var count int 

// Node represents the components of a binary search tree
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert will add a node to the tree
// The key to add should not be already in the tree
func (n *Node) Insert(k int) {
	if n.Key < k {
		// Move Right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		// Move Left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

// Search will take in a key value and return true if there is a node with that value
func (n *Node) Search(k int) bool {

	count++

	if n == nil {
		return false
	}

	if n.Key < k {
		// Move Right
		return n.Right.Search(k)
	} else if n.Key > k {
		// Move Left
		return n.Left.Search(k)
	}
	return true
}

func main() {
	// Tree is a Address Variable (because we will call as a pointer to edit the real variable)
	tree := &Node{Key: 100}

	fmt.Println(tree)

	tree.Insert(52)
	tree.Insert(203)
	tree.Insert(19)
	tree.Insert(76)
	tree.Insert(150)
	tree.Insert(310)
	tree.Insert(7)
	tree.Insert(24)
	tree.Insert(88)
	tree.Insert(276)

	fmt.Println(tree.Search(400))
	fmt.Printf("There is %d nodes into this tree", count)
}
