// package hash
package hash_test

import (
	"fmt"

	"github.com/eng618/go-eng/ds/hash"
)

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

// TODO: fix this test
// func TestNew(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		want *hash.Table
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := hash.New(); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("New() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestTable_Set(t *testing.T) {
// 	type fields struct {
// 		table map[uint64][]node
// 		hash  *maphash.Hash
// 	}
// 	type args struct {
// 		key   string
// 		value interface{}
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			h := &Table{
// 				table: tt.fields.table,
// 				hash:  tt.fields.hash,
// 			}
// 			h.Set(tt.args.key, tt.args.value)
// 		})
// 	}
// }

// func TestTable_Get(t *testing.T) {
// 	t.Parallel()

// 	type fields struct {
// 		table map[uint64][]node
// 		hash  *maphash.Hash
// 	}

// 	type args struct {
// 		key string
// 	}
// 	tests := []struct {
// 		name      string
// 		fields    fields
// 		args      args
// 		wantValue interface{}
// 		wantErr   bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		tt := tt
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()
// 			gotValue, err := h.Get(tt.args.key)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Table.Get() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(gotValue, tt.wantValue) {
// 				t.Errorf("Table.Get() = %v, want %v", gotValue, tt.wantValue)
// 			}
// 		})
// 	}
// }

// func TestTable_Keys(t *testing.T) {
// 	type fields struct {
// 		table map[uint64][]node
// 		hash  *maphash.Hash
// 	}
// 	tests := []struct {
// 		name     string
// 		fields   fields
// 		wantKeys []string
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			h := &Table{
// 				table: tt.fields.table,
// 				hash:  tt.fields.hash,
// 			}
// 			if gotKeys := h.Keys(); !reflect.DeepEqual(gotKeys, tt.wantKeys) {
// 				t.Errorf("Table.Keys() = %v, want %v", gotKeys, tt.wantKeys)
// 			}
// 		})
// 	}
// }

// func TestTable_Values(t *testing.T) {
// 	type fields struct {
// 		table map[uint64][]node
// 		hash  *maphash.Hash
// 	}
// 	tests := []struct {
// 		name       string
// 		fields     fields
// 		wantValues []interface{}
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			h := &Table{
// 				table: tt.fields.table,
// 				hash:  tt.fields.hash,
// 			}
// 			if gotValues := h.Values(); !reflect.DeepEqual(gotValues, tt.wantValues) {
// 				t.Errorf("Table.Values() = %v, want %v", gotValues, tt.wantValues)
// 			}
// 		})
// 	}
// }

// func TestTable_Print(t *testing.T) {
// 	type fields struct {
// 		table map[uint64][]node
// 		hash  *maphash.Hash
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			h := &hash.Table{
// 				table: tt.fields.table,
// 				hash:  tt.fields.hash,
// 			}
// 			h.Print()
// 		})
// 	}
// }
