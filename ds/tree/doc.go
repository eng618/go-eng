/*
Package tree provides a simple implementation of a binary tree data structure.

A binary tree is a hierarchical data structure in which each node has at most two children,
referred to as the left child and the right child. This implementation follows the binary
search tree (BST) property where:
- All values in the left subtree are less than the node's value
- All values in the right subtree are greater than or equal to the node's value

Types:

- Node: Represents a single node in the binary tree with a value and pointers to left and right children.
- Tree: The main structure representing the binary tree with a root node.

Functions:

- New: Creates and returns a new empty binary tree.
- NewWithRoot: Creates and returns a new binary tree with a specified root value.
- Insert: Adds a new value to the binary tree following BST properties.
- Search: Looks for a value in the binary tree and returns true if found.
- Delete: Removes a value from the binary tree and returns an error if not found.

Example Usage:

	// Create a new empty tree
	tr := tree.New()

	// Insert values
	tr.Insert(10)
	tr.Insert(5)
	tr.Insert(15)
	tr.Insert(3)
	tr.Insert(7)

	// Search for a value
	found := tr.Search(7) // returns true
	notFound := tr.Search(100) // returns false

	// Delete a value
	err := tr.Delete(5)
	if err != nil {
		// handle error
	}

For more information about binary trees, see:
https://www.geeksforgeeks.org/binary-tree-set-1-introduction/
*/
package tree
