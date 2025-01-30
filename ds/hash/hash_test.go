package hash_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/eng618/go-eng/ds/hash"
)

// ExamplePrint demonstrates the usage of the hash package by creating a new hash table,
// setting key-value pairs, and retrieving values. The expected output shows the state
// of the hash table after each operation.
func ExamplePrint() {
	h := hash.New()
	h.Print()
	h.Set("cool", "HashTables are cool")
	h.Set("best", "HashTables are the best")

	fmt.Println(h.Get("cool"))
	fmt.Println(h.Get("best"))

	// Output:
	// map[]
	// HashTables are cool <nil>
	// HashTables are the best <nil>
}

// TestNew tests the New function of the hash package.
// It iterates over a list of test cases, each containing a name and the expected result.
// For each test case, it runs the New function and compares the result with the expected value using reflect.DeepEqual.
// If the result does not match the expected value, it reports an error.
func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *hash.Table
	}{
		{name: "Empty hash table", want: hash.New()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hash.New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestTable_Set tests the Set method of the hash table implementation.
// It verifies that values can be correctly set and retrieved from the hash table.
// The test sets two key-value pairs and checks if the values can be retrieved correctly.
func TestTable_Set(t *testing.T) {
	h := hash.New()
	h.Set("key1", "value1")
	h.Set("key2", "value2")

	if got, _ := h.Get("key1"); got != "value1" {
		t.Errorf("Set() = %v, want %v", got, "value1")
	}
	if got, _ := h.Get("key2"); got != "value2" {
		t.Errorf("Set() = %v, want %v", got, "value2")
	}
}

// TestTable_Get tests the Get method of the hash table implementation.
// It initializes a new hash table, sets a key-value pair, and defines
// a set of test cases to verify the behavior of the Get method.
// The test cases include scenarios for an existing key and a non-existing key.
// It checks if the returned value and error match the expected results.
func TestTable_Get(t *testing.T) {
	h := hash.New()
	h.Set("key1", "value1")

	tests := []struct {
		name    string
		key     string
		want    interface{}
		wantErr bool
	}{
		{"Existing key", "key1", "value1", false},
		{"Non-existing key", "key2", nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := h.Get(tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestTable_Keys tests the Keys method of the hash table.
// It verifies that the Keys method returns the correct slice of keys
// that have been set in the hash table.
func TestTable_Keys(t *testing.T) {
	h := hash.New()
	h.Set("key1", "value1")
	h.Set("key2", "value2")

	wantKeys := []string{"key1", "key2"}
	if gotKeys := h.Keys(); !reflect.DeepEqual(gotKeys, wantKeys) {
		t.Errorf("Keys() = %v, want %v", gotKeys, wantKeys)
	}
}

// TestTable_Values tests the Values method of the hash table implementation.
// It verifies that the Values method returns the correct slice of values
// that were previously set in the hash table.
func TestTable_Values(t *testing.T) {
	h := hash.New()
	h.Set("key1", "value1")
	h.Set("key2", "value2")

	wantValues := []interface{}{"value1", "value2"}
	if gotValues := h.Values(); !reflect.DeepEqual(gotValues, wantValues) {
		t.Errorf("Values() = %v, want %v", gotValues, wantValues)
	}
}

// TestTable_Print tests the Print method of the hash table.
// It creates a new hash table, sets two key-value pairs, and then prints the hash table.
func TestTable_Print(_ *testing.T) {
	h := hash.New()
	h.Set("key1", "value1")
	h.Set("key2", "value2")

	h.Print()
}

// TestTable_EdgeCases tests edge cases for the hash table implementation.
func TestTable_EdgeCases(t *testing.T) {
	h := hash.New()

	// Test setting and getting an empty key
	h.Set("", "emptyKey")
	if got, _ := h.Get(""); got != "emptyKey" {
		t.Errorf("Get() with empty key = %v, want %v", got, "emptyKey")
	}

	// Test setting and getting a very long key
	longKey := string(make([]byte, 1000))
	h.Set(longKey, "longKey")
	if got, _ := h.Get(longKey); got != "longKey" {
		t.Errorf("Get() with long key = %v, want %v", got, "longKey")
	}

	// Test setting and getting a nil value
	h.Set("nilValue", nil)
	if got, _ := h.Get("nilValue"); got != nil {
		t.Errorf("Get() with nil value = %v, want %v", got, nil)
	}
}

// BenchmarkTable_Set benchmarks the Set method of the hash table.
func BenchmarkTable_Set(b *testing.B) {
	h := hash.New()
	for i := 0; i < b.N; i++ {
		h.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}
}

// BenchmarkTable_Get benchmarks the Get method of the hash table.
func BenchmarkTable_Get(b *testing.B) {
	h := hash.New()
	for i := 0; i < 1000; i++ {
		h.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := h.Get(fmt.Sprintf("key%d", i%1000)); err != nil {
			b.Errorf("Get() error = %v", err)
		}
	}
}

// BenchmarkTable_Keys benchmarks the Keys method of the hash table.
func BenchmarkTable_Keys(b *testing.B) {
	h := hash.New()
	for i := 0; i < 1000; i++ {
		h.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Keys()
	}
}

// BenchmarkTable_Values benchmarks the Values method of the hash table.
func BenchmarkTable_Values(b *testing.B) {
	h := hash.New()
	for i := 0; i < 1000; i++ {
		h.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Values()
	}
}
