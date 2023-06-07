/*
Package list is used to demonstrates a linked list data structure.

# Linked List Data Structure

Some good code optimizations for linked lists are to have head AND tail pointers. This will allow you to add items to the end of the list in O(1) constant time rather than O(n) linear time. Another best practice to help traverse the list is to add a previous pointer to each node of the linked list.

I started with this [code source](https://divyanshushekhar.com/golang-linked-list/) and expanded on it to test and document the code using Go's best practices. Also while writing tests, I found a bug in the code that it was not properly deleting the final element of a list.
*/
package list
