// Package tree provides a simple implementation of a binary tree data structure.
package tree

import (
	"errors"
)

// Node represents a node in the binary tree.
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// Tree represents a binary tree structure.
type Tree struct {
	Root *Node
}

// New creates and returns a new empty binary tree.
func New() *Tree {
	return &Tree{Root: nil}
}

// NewWithRoot creates and returns a new binary tree with the given root value.
func NewWithRoot(value int) *Tree {
	return &Tree{Root: &Node{Value: value}}
}

// Insert adds a new value to the binary tree.
// The tree is constructed as a binary search tree where:
// - values less than the current node go to the left
// - values greater than or equal to the current node go to the right
func (t *Tree) Insert(value int) {
	if t.Root == nil {
		t.Root = &Node{Value: value}
		return
	}
	t.Root.insert(value)
}

func (n *Node) insert(value int) {
	if value < n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: value}
		} else {
			n.Left.insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Value: value}
		} else {
			n.Right.insert(value)
		}
	}
}

// Search looks for a value in the binary tree.
// Returns true if the value is found, false otherwise.
func (t *Tree) Search(value int) bool {
	if t.Root == nil {
		return false
	}
	return t.Root.search(value)
}

func (n *Node) search(value int) bool {
	if n == nil {
		return false
	}
	if n.Value == value {
		return true
	}
	if value < n.Value {
		return n.Left.search(value)
	}
	return n.Right.search(value)
}

// Delete removes a value from the binary tree.
// Returns an error if the tree is empty or the value is not found.
func (t *Tree) Delete(value int) error {
	if t.Root == nil {
		return errors.New("tree is empty")
	}
	
	var deleted bool
	t.Root, deleted = t.Root.delete(value)
	if !deleted {
		return errors.New("value not found")
	}
	return nil
}

func (n *Node) delete(value int) (*Node, bool) {
	if n == nil {
		return nil, false
	}
	
	var deleted bool
	if value < n.Value {
		n.Left, deleted = n.Left.delete(value)
	} else if value > n.Value {
		n.Right, deleted = n.Right.delete(value)
	} else {
		deleted = true
		// Node with only one child or no child
		if n.Left == nil {
			return n.Right, deleted
		} else if n.Right == nil {
			return n.Left, deleted
		}
		
		// Node with two children: Get the inorder successor (smallest in the right subtree)
		minNode := n.Right.findMin()
		n.Value = minNode.Value
		n.Right, _ = n.Right.delete(minNode.Value)
	}
	
	return n, deleted
}

func (n *Node) findMin() *Node {
	current := n
	for current.Left != nil {
		current = current.Left
	}
	return current
}
