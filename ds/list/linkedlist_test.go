package list_test

import (
	"fmt"
	"testing"

	"github.com/eng618/go-eng/ds/list"
)

// testdata initializes a LinkedList with a predefined set of elements and returns
// the LinkedList along with a slice of the same elements. The elements include
// integers, booleans, and a string.
//
// Returns:
//   - list.LinkedList: The initialized LinkedList containing the predefined elements.
//   - []interface{}: A slice containing the same elements as the LinkedList.
func testdata() (*list.LinkedList, []interface{}) {
	ll := list.NewLinkedList()
	nx := []interface{}{
		10,
		20,
		30,
		40,
		50,
		true,
		false,
		"strings work too",
	}

	for _, v := range nx {
		ll.PushBack(v)
	}

	return &ll, nx
}

// ----------------------------------------------------------------
// Examples

func Example() {
	list := list.NewLinkedList()

	list.PushBack(20)
	list.PushBack(30)
	list.PushBack(40)
	list.PushBack(50)
	list.PushBack(70)

	fmt.Println("Length =", list.Length())

	list.Display()

	if err := list.Delete(2); err != nil {
		fmt.Println(err)
	}

	list.Reverse()

	fmt.Println("Length =", list.Length())

	list.Display()

	front, _ := list.PeekFront()
	back, _ := list.PeekBack()

	fmt.Println("Front =", front)
	fmt.Println("Back =", back)

	// Output:
	// Length = 5
	// 20 -> 30 -> 40 -> 50 -> 70 ->
	// Length = 4
	// 70 -> 50 -> 30 -> 20 ->
	// Front = 70
	// Back = 20
}

func ExampleLinkedList() {
	ll := list.NewLinkedList()
	ll.PushBack(1)
	ll.PushBack(2)
	ll.PushFront(0)
	ll.Display()
	// Output: 0 -> 1 -> 2 ->
}

// ----------------------------------------------------------------
// Tests

func TestPeekFront(t *testing.T) {
	ll, _ := testdata()
	front, err := ll.PeekFront()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if front != 10 {
		t.Errorf("expected 10, got %v", front)
	}
}

func TestPeekBack(t *testing.T) {
	ll, _ := testdata()
	back, err := ll.PeekBack()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if back != "strings work too" {
		t.Errorf("expected 'strings work too', got %v", back)
	}
}

func TestLength(t *testing.T) {
	ll, data := testdata()
	if ll.Length() != len(data) {
		t.Errorf("expected length %d, got %d", len(data), ll.Length())
	}
}

func TestPushBack(t *testing.T) {
	ll := list.NewLinkedList()
	ll.PushBack(1)
	if ll.Length() != 1 {
		t.Errorf("expected length 1, got %d", ll.Length())
	}
	if ll.Head() != 1 {
		t.Errorf("expected head 1, got %v", ll.Head())
	}
	if ll.Tail() != 1 {
		t.Errorf("expected tail 1, got %v", ll.Tail())
	}
}

func TestPushFront(t *testing.T) {
	ll := list.NewLinkedList()
	ll.PushFront(1)
	if ll.Length() != 1 {
		t.Errorf("expected length 1, got %d", ll.Length())
	}
	if ll.Head() != 1 {
		t.Errorf("expected head 1, got %v", ll.Head())
	}
	if ll.Tail() != 1 {
		t.Errorf("expected tail 1, got %v", ll.Tail())
	}
}

func TestRemove(t *testing.T) {
	ll, _ := testdata()
	err := ll.Remove(10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ll.Length() != 7 {
		t.Errorf("expected length 7, got %d", ll.Length())
	}
	if ll.Head() != 20 {
		t.Errorf("expected head 20, got %v", ll.Head())
	}
}

func TestDelete(t *testing.T) {
	ll, _ := testdata()
	err := ll.Delete(0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ll.Length() != 7 {
		t.Errorf("expected length 7, got %d", ll.Length())
	}
	if ll.Head() != 20 {
		t.Errorf("expected head 20, got %v", ll.Head())
	}
}

func TestReverse(t *testing.T) {
	ll, _ := testdata()
	ll.Reverse()
	if ll.Head() != "strings work too" {
		t.Errorf("expected head 'strings work too', got %v", ll.Head())
	}
	if ll.Tail() != 10 {
		t.Errorf("expected tail 10, got %v", ll.Tail())
	}
}

// Edge case tests

func TestDeleteEmptyList(t *testing.T) {
	ll := list.NewLinkedList()
	err := ll.Delete(0)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestDeleteOutOfRange(t *testing.T) {
	ll := list.NewLinkedList()
	ll.PushBack(1)
	err := ll.Delete(1)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestRemoveEmptyList(t *testing.T) {
	ll := list.NewLinkedList()
	err := ll.Remove(1)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestPeekFrontEmptyList(t *testing.T) {
	ll := list.NewLinkedList()
	_, err := ll.PeekFront()
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestPeekBackEmptyList(t *testing.T) {
	ll := list.NewLinkedList()
	_, err := ll.PeekBack()
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestNewSeeded(t *testing.T) {
	ll := list.NewSeeded(100)
	if ll.Length() != 1 {
		t.Errorf("expected length 1, got %d", ll.Length())
	}
	if ll.Head() != 100 {
		t.Errorf("expected head 100, got %v", ll.Head())
	}
	if ll.Tail() != 100 {
		t.Errorf("expected tail 100, got %v", ll.Tail())
	}
}

func TestNewSeededWithString(t *testing.T) {
	ll := list.NewSeeded("seed")
	if ll.Length() != 1 {
		t.Errorf("expected length 1, got %d", ll.Length())
	}
	if ll.Head() != "seed" {
		t.Errorf("expected head 'seed', got %v", ll.Head())
	}
	if ll.Tail() != "seed" {
		t.Errorf("expected tail 'seed', got %v", ll.Tail())
	}
}

func TestNewSeededWithBool(t *testing.T) {
	ll := list.NewSeeded(true)
	if ll.Length() != 1 {
		t.Errorf("expected length 1, got %d", ll.Length())
	}
	if ll.Head() != true {
		t.Errorf("expected head true, got %v", ll.Head())
	}
	if ll.Tail() != true {
		t.Errorf("expected tail true, got %v", ll.Tail())
	}
}

// ----------------------------------------------------------------
// Benchmarks

func BenchmarkPushBack(b *testing.B) {
	ll := list.NewLinkedList()
	for i := 0; i < b.N; i++ {
		ll.PushBack(i)
	}
}

func BenchmarkPushFront(b *testing.B) {
	ll := list.NewLinkedList()
	for i := 0; i < b.N; i++ {
		ll.PushFront(i)
	}
}
