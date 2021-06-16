// package hash
package hash

import (
	"errors"
	"fmt"
	"hash/maphash"
)

type HashTable struct {
	table map[uint64][]node
}

type node struct {
	key   string
	value interface{}
}

var h maphash.Hash

// New creates a HashTable with a specified size
func New() *HashTable {
	return &HashTable{
		table: make(map[uint64][]node),
	}
}

// hash is a implementation using hash/maphash from the go standard library
func hash(key string) (hash uint64) {
	_, _ = h.WriteString(key)
	hash = h.Sum64()
	h.Reset()
	return
}

// Set allows you to set the value for a new item in a HashTable
func (h *HashTable) Set(key string, value interface{}) {
	k := hash(key)
	if h.table[k] == nil {
		h.table[k] = make([]node, 0)
	}
	h.table[k] = append(h.table[k], node{key: key, value: value})
}

// Get returns the value for the specified key.
func (h *HashTable) Get(key string) (value interface{}, err error) {
	k := hash(key)
	v := h.table[k]

	for _, xValue := range v {
		if xValue.key == key {
			return xValue.value, nil
		}
	}

	return nil, errors.New("there is no value associated with this key")
}

// Keys returns a slice of all the keys in the HashTable.
func (h *HashTable) Keys() (keys []string) {
	for _, v := range h.table {
		for _, xValue := range v {
			keys = append(keys, xValue.key)
		}
	}
	return
}

// Values returns a slice of all the keys in the HashTable.
func (h *HashTable) Values() (values []interface{}) {
	for _, v := range h.table {
		for _, xValue := range v {
			values = append(values, xValue.value)
		}
	}
	return
}

func (h *HashTable) Print() {
	fmt.Println(h.table)
}
