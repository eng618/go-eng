package db

import (
	"fmt"
	"time"
)

// InMemDB is a simple in memory database.
type InMemDB struct {
	data map[string]map[time.Time]interface{}
}

// NewDatabase instantiates a new instance of a database.
func NewDatabase() *InMemDB {
	return &InMemDB{
		data: make(map[string]map[time.Time]interface{}),
	}
}

// Set creates a new entry into an existing database.
func (db *InMemDB) Set(key string, value interface{}) time.Time {
	t := time.Now()
	if v, ok := db.data[key]; ok {
		// Here we need to update v
		v[t] = value
		return t
	}

	newVal := make(map[time.Time]interface{})
	newVal[t] = value

	db.data[key] = newVal
	return t
}

// Get attempts to retrieve the supplied key from an existing database.
func (db *InMemDB) Get(key string) interface{} {
	if v, ok := db.data[key]; ok {
		var ot time.Time
		var finalValue interface{}
		for t, val := range v {
			if t.After(ot) {
				ot = t
				finalValue = val
			}
		}

		return finalValue
	}

	fmt.Println("Value is not in database")
	return nil
}

// Get attempts to retrieve the supplied key from an existing database.
func (db *InMemDB) GetForTime(key string, t time.Time) (interface{}, bool) {
	if v, ok := db.data[key]; ok {
		if finalValue, ok := v[t]; ok {
			return finalValue, true
		}
	}

	fmt.Println("No value for:", key, "at:", t, "found in database.")
	return nil, false
}

// Print is a helper method to print each entry in a database to it's own line.
func (db *InMemDB) Print() {
	for k, v := range db.data {
		fmt.Printf("%v: %v\n", k, v)
	}
}
