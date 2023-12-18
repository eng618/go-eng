package db_test

import (
	"fmt"
	"time"

	"github.com/eng618/go-eng/interview/db"
)

func Example() {
	db := db.NewDatabase()

	db.Set("foo", "bar")
	fmt.Println("just set foo to:", db.Get("foo"))

	db.Set("foo", "baz")
	fmt.Println("the latest foo is:", db.Get("foo"))

	// DB can set ints also
	db.Set("age", 36)
	age := db.Get("age")
	fmt.Println("the age in database is:", age)

	// You can use the common ok idiom to get a key with a particular time stamp.
	if val, ok := db.GetForTime("foo", time.Now()); ok {
		fmt.Println(val)
	}

	db.Print()
}
