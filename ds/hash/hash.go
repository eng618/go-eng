// package hash
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

// New creates a hash table.
func New() *Table {
	var hash maphash.Hash

	return &Table{
		table: make(map[uint64][]node),
		hash:  &hash,
	}
}

// hash is an implementation using hash/maphash from the go standard library.
func hash(key string, h *maphash.Hash) (hash uint64) {
	_, _ = h.WriteString(key)
	hash = h.Sum64()
	h.Reset()

	return
}

// Set allows you to set the value for a new item in a hash table.
func (h *Table) Set(key string, value interface{}) {
	k := hash(key, h.hash)
	if h.table[k] == nil {
		h.table[k] = make([]node, 0)
	}

	h.table[k] = append(h.table[k], node{key: key, value: value})
}

// Get returns the value for the specified key.
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

// Keys returns a slice of all the keys in the HashTable.
func (h *Table) Keys() (keys []string) {
	for _, v := range h.table {
		for _, xValue := range v {
			keys = append(keys, xValue.key)
		}
	}

	return
}

// Values returns a slice of all the keys in the HashTable.
func (h *Table) Values() (values []interface{}) {
	for _, v := range h.table {
		for _, xValue := range v {
			values = append(values, xValue.value)
		}
	}

	return
}

func (h *Table) Print() {
	fmt.Println(h.table)
}
