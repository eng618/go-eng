package db_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/eng618/go-eng/interview/db"
)

func Example() {
	db := db.NewDatabase()

	t := db.Set("foo", "bar")
	fmt.Println("just set foo to:", db.Get("foo"))

	db.Set("foo", "baz")
	fmt.Println("the latest foo is:", db.Get("foo"))

	// You can use the common ok idiom to get a key with a particular time stamp.
	if val, ok := db.GetForTime("foo", t); ok {
		fmt.Println("the foo at", t, ":", val)
	}

	db.Set("age", 36)
	age := db.Get("age")
	fmt.Println("the age in database is:", age)

	db.Print()
}

func getTestDatabase() *db.InMemDB {
	db := db.NewDatabase()
	db.Set("foo", "bar")
	db.Set("hello", "world")
	return db
}

func TestNewDatabase(t *testing.T) {
	if got := db.NewDatabase(); got == nil {
		t.Errorf("NewDatabase() = %v, wanted not nil", got)
	}
}

func TestInMemDB_Set(t *testing.T) {
	db := getTestDatabase()
	if got := db.Get("foo"); got != "bar" {
		t.Error("Failed to get foo")
	}
}

func TestInMemDB_Set_multipleValuesForKey(t *testing.T) {
	db := getTestDatabase()
	if got := db.Get("foo"); got != "bar" {
		t.Error("Failed to get foo")
	}
}

func TestInMemDB_Get_basic(t *testing.T) {
	db := getTestDatabase()
	// foo originally set in getTestDatabase(), set an additional value
	db.Set("foo", "baz")
	if got := db.Get("foo"); got != "baz" {
		t.Error("Failed to get foo")
	}
}

func TestInMemDB_Get_noValue(t *testing.T) {
	db := getTestDatabase()
	if got := db.Get("none"); got != nil {
		t.Error("Failed to get foo")
	}
}

func TestInMemDB_GetForTime(t *testing.T) {
	db := getTestDatabase()
	setTime := db.Set("test", "testValue")

	if v, ok := db.GetForTime("test", setTime); ok {
		if v != "testValue" {
			t.Error("Failed, returned the wrong value.")
		}
	}
}

func TestInMemDB_GetForTime_noValue(t *testing.T) {
	db := getTestDatabase()
	if v, ok := db.GetForTime("foo", time.Now()); ok && v != nil {
		t.Error("this should be nil")
	}
}

func TestInMemDB_Print(_ *testing.T) {
	db := getTestDatabase()
	db.Print()
}
