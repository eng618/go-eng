// Package hashtable provides a basic implementation of a hash table data structure.
// It allows for insertion, retrieval, and deletion of key-value pairs using
// a hash function to compute the index for each key. The package also provides
// methods to retrieve all keys and values, as well as a method to print the
// entire hash table.
//
// The hash table is implemented using a map with uint64 keys and slices of nodes
// as values. Each node contains a string key and an interface{} value. The package
// uses the maphash package to compute hash values for keys.
//
// Example usage:
//
//	table := hashtable.New()
//	table.Set("key1", "value1")
//	value, err := table.Get("key1")
//	if err != nil {
//	    fmt.Println("Error:", err)
//	} else {
//	    fmt.Println("Value:", value)
//	}
//
//	keys := table.Keys()
//	values := table.Values()
//	table.Print()
//
// The package provides the following methods:
//
//   - New: Creates and returns a new instance of Table.
//   - Set: Inserts or updates the value associated with a given key.
//   - Get: Retrieves the value associated with a given key.
//   - Keys: Returns a slice of all keys in the hash table.
//   - Values: Returns a slice of all values in the hash table.
//   - Print: Outputs the contents of the hash table to the standard output.
package hashtable
