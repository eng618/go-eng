// Package hash provides a basic implementation of a hash table data structure.
// It allows for insertion, retrieval, and deletion of key-value pairs using
// a hash function to compute the index for each key. The package also provides
// methods to retrieve all keys and values, as well as a method to print the
// entire hash table.
package hash

import (
	"errors"
	"fmt"
	"hash/maphash"
)

// Table is a structured hash table.
type Table struct {
	table map[uint64][]node
	hash  *maphash.Hash
}

type node struct {
	key   string
	value interface{}
}

// New creates and returns a new instance of Table. It initializes the table
// with an empty map and a new maphash.Hash instance.
func New() *Table {
	var hash maphash.Hash

	return &Table{
		table: make(map[uint64][]node),
		hash:  &hash,
	}
}

// hash computes the hash value for a given key using the provided maphash.Hash.
// It writes the key to the hash, computes the hash value, resets the hash, and returns the computed value.
//
// Parameters:
//   - key: The string key to be hashed.
//   - h: A pointer to a maphash.Hash instance used for hashing.
//
// Returns:
//   - hash: The computed hash value as a uint64.
func hash(key string, h *maphash.Hash) (hash uint64) {
	_, _ = h.WriteString(key)
	hash = h.Sum64()
	h.Reset()

	return
}

// Set inserts or updates the value associated with the given key in the hash table.
// If the key already exists, its value is updated with the new value.
// If the key does not exist, a new key-value pair is added to the table.
//
// Parameters:
//
//	key: A string representing the key to be added or updated in the hash table.
//	value: An interface{} representing the value to be associated with the key.
func (h *Table) Set(key string, value interface{}) {
	k := hash(key, h.hash)
	if h.table[k] == nil {
		h.table[k] = make([]node, 0)
	}

	h.table[k] = append(h.table[k], node{key: key, value: value})
}

// Get retrieves the value associated with the given key from the hash table.
// It returns the value if the key is found, otherwise it returns an error indicating
// that there is no value associated with the key.
//
// Parameters:
//   - key: The key to search for in the hash table.
//
// Returns:
//   - value: The value associated with the key, if found.
//   - err: An error if the key is not found in the hash table.
func (h *Table) Get(key string) (value interface{}, err error) {
	k := hash(key, h.hash)
	v := h.table[k]

	for _, xValue := range v {
		if xValue.key == key {
			return xValue.value, nil
		}
	}

	return nil, errors.New("there is no value associated with this key")
}

// Keys returns a slice of all the keys present in the hash table.
// It iterates through each bucket in the table and collects the keys
// from each entry in the bucket.
func (h *Table) Keys() (keys []string) {
	for _, v := range h.table {
		for _, xValue := range v {
			keys = append(keys, xValue.key)
		}
	}

	return
}

// Values returns a slice containing all the values stored in the hash table.
// It iterates through each bucket in the table and appends each value to the
// resulting slice.
//
// Returns:
//
//	[]interface{} - A slice of all values in the hash table.
func (h *Table) Values() (values []interface{}) {
	for _, v := range h.table {
		for _, xValue := range v {
			values = append(values, xValue.value)
		}
	}

	return
}

// Print outputs the contents of the hash table to the standard output.
func (h *Table) Print() {
	fmt.Println(h.table)
}
