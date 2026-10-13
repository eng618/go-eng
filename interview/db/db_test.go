package db_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/eng618/go-eng/interview/db"
)

func ExampleNewDatabase() {
	database := db.NewDatabase()

	// Getting a key that was never set prints a notice and returns nil.
	fmt.Println(database.Get("missing"))
	// Output:
	// Value is not in database
	// <nil>
}

func TestDatabaseWorkflow(t *testing.T) {
	database := db.NewDatabase()

	setTime := database.Set("foo", "bar")
	if got := database.Get("foo"); got != "bar" {
		t.Errorf("Get(foo) = %v, want %v", got, "bar")
	}

	database.Set("foo", "baz")
	if got := database.Get("foo"); got != "baz" {
		t.Errorf("Get(foo) = %v, want %v", got, "baz")
	}

	// The common ok idiom retrieves a key with a particular time stamp.
	if val, ok := database.GetForTime("foo", setTime); !ok || val != "bar" {
		t.Errorf("GetForTime(foo, setTime) = (%v, %v), want (%v, %v)", val, ok, "bar", true)
	}

	database.Set("age", 36)
	if got := database.Get("age"); got != 36 {
		t.Errorf("Get(age) = %v, want %v", got, 36)
	}
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
