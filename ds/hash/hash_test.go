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
