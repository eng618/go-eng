package list_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/eng618/go-eng/ds/list"
)

// testdata is a helper function to create a base LinkedList for running tests against.
func testdata() (list.LinkedList, []interface{}) {
	ll := list.New()
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

	return ll, nx
}

func Example() {
	list := list.New()

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

// nolint:paralleltest
func TestLinkedList_Delete(t *testing.T) {
	type args struct {
		position int
	}

	ll, _ := testdata()
	tests := []struct {
		name string
		args args
	}{
		// {20, 30, 40, 50, true, false}
		{name: "non existent position", args: args{position: 100}},
		{name: "tail position", args: args{position: (ll.Length() - 1)}},
		{name: "head position", args: args{position: 0}},
		{name: "position 5", args: args{position: 5}},
		{name: "position 4", args: args{position: 4}},
		{name: "position 3", args: args{position: 3}},
		{name: "position 2", args: args{position: 2}},
		{name: "position 1", args: args{position: 1}},
		{name: "position 0", args: args{position: 0}},
		{name: "empty list", args: args{position: 40}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = ll.Delete(tt.args.position)
		})
	}
}

func TestLinkedList_Display(t *testing.T) {
	t.Parallel()
}

// nolint:paralleltest
func TestLinkedList_PeekBack(t *testing.T) {
	ll, nx := testdata()

	for i := 0; i < len(nx); i++ {
		td := nx[ll.Length()-1+i]

		t.Run(fmt.Sprintf("Back value: %v", td), func(t *testing.T) {
			got, err := ll.PeekBack()
			if err != nil {
				t.Errorf("LinkedList.Back() error = %v, wantErr %v", err, nil)

				return
			}

			if got != td {
				t.Errorf("LinkedList.Back() = %v, want %v", got, td)
			}

			if err := ll.Delete(ll.Length() - 1); err != nil {
				t.Errorf("LinkedList.Back() failed to delete item")
			}
		})
	}

	t.Run("Check back on empty list", func(t *testing.T) {
		got, err := ll.PeekBack()
		if err == nil {
			t.Errorf("LinkedList.Back() error = %v, wantErr %v",
				err, fmt.Errorf("Cannot find Back value in an empty linked list"))

			return
		}
		if got != nil {
			t.Errorf("LinkedList.Back() = %v, want 0", got)
		}
	})
}

// nolint:paralleltest
func TestLinkedList_PeekFront(t *testing.T) {
	ll, nx := testdata()

	for _, v := range nx {
		t.Run(fmt.Sprintf("Front value: %v", v), func(t *testing.T) {
			got, err := ll.PeekFront()
			if err != nil {
				t.Errorf("LinkedList.Front() error = %v, wantErr %v", err, nil)

				return
			}
			if got != v {
				t.Errorf("LinkedList.Front() = %v, want %v", got, v)
			}
		})

		_ = ll.Delete(0)
	}

	t.Run("Check head on empty list", func(t *testing.T) {
		got, err := ll.PeekFront()
		if err == nil {
			t.Errorf("LinkedList.Front() error = %v, wantErr %v",
				err, fmt.Errorf("Cannot Find Front Value in an Empty linked list"))

			return
		}
		if got != nil {
			t.Errorf("LinkedList.Front() = %v, want 0", got)
		}
	})
}

func TestLinkedList_PushBack(t *testing.T) {
	t.Parallel()

	ll, nx := testdata()
	emptyList := list.New()

	tests := []struct {
		name string
		list list.LinkedList
		args []interface{}
		want interface{}
	}{
		{"Test tail", ll, []interface{}{nx[1], nx[2]}, nx[2]},
		{"Test tail 2", ll, []interface{}{nx[4], nx[3]}, nx[3]},
		{"Empty list", emptyList, []interface{}{nx[0]}, nx[0]},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			l := tt.list
			for _, v := range tt.args {
				l.PushBack(v)
			}
			if !reflect.DeepEqual(l.Tail(), tt.want) {
				t.Error("Expected:", tt.want, "Got:", l.Tail())
			}
		})
	}
}

func TestLinkedList_PushFront(t *testing.T) {
	t.Parallel()

	ll, nx := testdata()

	tests := []struct {
		name   string
		list   list.LinkedList
		expect interface{}
	}{
		{"Test head", ll, nx[0]},
		{"Test head 1", ll, nx[1]},
		{"Test head 2", ll, nx[2]},
		{"Test head 3", ll, nx[3]},
		{"Empty list", list.New(), nx[0]},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.list.PushFront(tt.expect)
			if tt.expect != tt.list.Head() {
				t.Error("Expected:", tt.expect, "Got:", tt.list.Head())
			}
		})
	}
}

// nolint:paralleltest
func TestLinkedList_Remove(t *testing.T) {
	type args struct {
		data interface{}
	}

	ll, nx := testdata()
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Not Found", args: args{data: "not in list"}, wantErr: true},
		{name: "Remove 6", args: args{data: nx[6]}, wantErr: false},
		{name: "Remove 4", args: args{data: nx[4]}, wantErr: false},
		{name: "Remove 2", args: args{data: nx[2]}, wantErr: false},
		{name: "Remove 0", args: args{data: nx[0]}, wantErr: false},
		{name: "Remove 1", args: args{data: nx[1]}, wantErr: false},
		{name: "Remove 3", args: args{data: nx[3]}, wantErr: false},
		{name: "Remove 5", args: args{data: nx[5]}, wantErr: false},
		{name: "Remove 7", args: args{data: nx[7]}, wantErr: false},
		{name: "Empty list", args: args{data: nx[7]}, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ll.Remove(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("LinkedList.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLinkedList_Reverse(t *testing.T) {
	t.Parallel()

	ll, nx := testdata()

	t.Run("Simple reverse", func(t *testing.T) {
		t.Parallel()
		ll.Reverse()
		if ll.Head() != nx[len(nx)-1] || ll.Tail() != nx[0] {
			t.Errorf("LinkedList.Reverse() values were not reversed")
		}
	})
}

func TestLinkedList_Length(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		ll   list.LinkedList
		want int
	}{
		{"Empty list", list.New(), 0},
		{"Length of 2", list.NewSeeded("test"), 1},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.ll.Length(); got != tt.want {
				t.Errorf("LinkedList.Length() = %v, want %v", got, tt.want)
			}
		})
	}
}
