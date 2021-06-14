package list

import (
	"fmt"
	"reflect"
	"testing"
)

func Example() {
	list := New()

	list.PushBack(20)
	list.PushBack(30)
	list.PushBack(40)
	list.PushBack(50)
	list.PushBack(70)

	fmt.Println("Length =", list.Len())

	list.Display()

	list.Delete(2)

	list.Reverse()

	fmt.Println("Length =", list.Len())

	list.Display()

	front, _ := list.Front()
	back, _ := list.Back()
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

func TestLinkedList_Len(t *testing.T) {
	type fields struct {
		Length int
		head   *node
		tail   *node
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"Empty list", fields{}, 0},
		{"Len of 2", fields{Length: 2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				Length: tt.fields.Length,
				head:   tt.fields.head,
				tail:   tt.fields.tail,
			}
			if got := l.Len(); got != tt.want {
				t.Errorf("LinkedList.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_PushFront(t *testing.T) {
	ll, nx := testdata()

	tests := []struct {
		name   string
		list   LinkedList
		expect interface{}
	}{
		{"Test head", ll, nx[0]},
		{"Test head 1", ll, nx[1]},
		{"Test head 2", ll, nx[2]},
		{"Test head 3", ll, nx[3]},
		{"Empty list", New(), nx[0]},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.list.PushFront(tt.expect)
			if tt.expect != tt.list.head.data {
				t.Error("Expected:", tt.expect, "Got:", tt.list.head.data)
			}
		})

	}
}

func TestLinkedList_PushBack(t *testing.T) {
	ll, nx := testdata()
	emptyList := New()

	tests := []struct {
		name string
		list LinkedList
		args []interface{}
		want *node
	}{
		{"Test tail", ll, []interface{}{nx[1], nx[2]}, &node{data: nx[2]}},
		{"Test tail 2", ll, []interface{}{nx[4], nx[3]}, &node{data: nx[3]}},
		{"Empty list", emptyList, []interface{}{nx[0]}, &node{data: nx[0]}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.list
			for _, v := range tt.args {
				l.PushBack(v)
			}
			if !reflect.DeepEqual(l.tail, tt.want) {
				t.Error("Expected:", tt.want, "Got:", l.tail)
			}
		})
	}
}

func TestLinkedList_Delete(t *testing.T) {
	ll, _ := testdata()

	type args struct {
		position int
	}
	tests := []struct {
		name string
		args args
	}{
		// {20, 30, 40, 50, true, false}
		{name: "non existent position", args: args{position: 100}},
		{name: "tail position", args: args{position: (ll.Length - 1)}},
		{name: "head position", args: args{position: 0}},
		{name: "postion 5", args: args{position: 5}},
		{name: "postion 4", args: args{position: 4}},
		{name: "postion 3", args: args{position: 3}},
		{name: "postion 2", args: args{position: 2}},
		{name: "postion 1", args: args{position: 1}},
		{name: "postion 0", args: args{position: 0}},
		{name: "empty list", args: args{position: 40}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll.Delete(tt.args.position)
		})
	}
}

func TestLinkedList_Front(t *testing.T) {
	ll, nx := testdata()

	for _, v := range nx {
		t.Run(fmt.Sprintf("Front value: %v", v), func(t *testing.T) {
			got, err := ll.Front()
			if err != nil {
				t.Errorf("LinkedList.Front() error = %v, wantErr %v", err, nil)
				return
			}
			if got != v {
				t.Errorf("LinkedList.Front() = %v, want %v", got, v)
			}
		})
		ll.Delete(0)
	}

	t.Run("Check head on empty list", func(t *testing.T) {
		got, err := ll.Front()
		if err == nil {
			t.Errorf("LinkedList.Front() error = %v, wantErr %v", err, fmt.Errorf("Cannot Find Front Value in an Empty linked list"))
			return
		}
		if got != 0 {
			t.Errorf("LinkedList.Front() = %v, want 0", got)
		}
	})
}

func TestLinkedList_Back(t *testing.T) {
	ll, nx := testdata()

	for i := 0; i < len(nx); i++ {
		td := ll.tail.data

		t.Run(fmt.Sprintf("Back value: %v", td), func(t *testing.T) {
			got, err := ll.Back()
			if err != nil {
				t.Errorf("LinkedList.Back() error = %v, wantErr %v", err, nil)
				return
			}
			if got != td {
				t.Errorf("LinkedList.Back() = %v, want %v", got, td)
			}
			ll.Delete(ll.Length - 1)
		})
	}

	t.Run("Check back on empty list", func(t *testing.T) {
		got, err := ll.Back()
		if err == nil {
			t.Errorf("LinkedList.Back() error = %v, wantErr %v", err, fmt.Errorf("Cannot find Back value in an empty linked list"))
			return
		}
		if got != 0 {
			t.Errorf("LinkedList.Back() = %v, want 0", got)
		}
	})
}

func TestLinkedList_Reverse(t *testing.T) {
	ll, nx := testdata()

	t.Run("Simple reverse", func(t *testing.T) {
		ll.Reverse()
		if ll.head.data != nx[len(nx)-1] || ll.tail.data != nx[0] {
			t.Errorf("LinkedList.Reverse() values were not reversed")
		}
	})
}

// testdata is a helper function to create a base LinkedList for running tests against.
func testdata() (LinkedList, []interface{}) {
	ll := New()
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
