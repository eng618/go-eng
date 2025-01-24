/*
Package stack provides a simple implementation of the stack data structure.

A stack is a linear data structure that follows the Last In, First Out (LIFO) principle.
The package includes methods for creating a new stack, pushing items onto the stack,
and popping items off the stack. It also supports concurrent access with a read-write mutex.

Types:

- Item: The type of elements the stack holds.
- Stack: The main structure representing the stack.

Functions:

- New: Creates and returns a new instance of Stack.
- NewSeeded: Creates a new Stack instance pre-seeded with the provided items.
- Push: Adds an item to the top of the stack.
- Pop: Removes and returns the top item from the stack.

Concurrency:

The stack is safe for concurrent use, with a read-write mutex protecting the internal slice of items.
*/
package stack
