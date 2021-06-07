package list

import (
	"fmt"
	"testing"
)

func Example() {
	list := LinkedList{}

	node1 := &Node{Data: 20}
	node2 := &Node{Data: 30}
	node3 := &Node{Data: 40}
	node4 := &Node{Data: 50}
	node5 := &Node{Data: 70}

	list.PushBack(node1)
	list.PushBack(node2)
	list.PushBack(node3)
	list.PushBack(node4)
	list.PushBack(node5)

	fmt.Println("Length =", list.Len())

	list.Display()

	list.Delete(40)

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
		length int
		head   *Node
		tail   *Node
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"Empty list", fields{}, 0},
		{"Len of 2", fields{length: 2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				length: tt.fields.length,
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
	ll, nx := testData()

	tests := []struct {
		name   string
		list   LinkedList
		args   *Node
		expect *Node
	}{
		{"Test head", ll, nx[0], nx[0]},
		{"Test head 1", ll, nx[1], nx[1]},
		{"Test head 2", ll, nx[2], nx[2]},
		{"Test head 3", ll, nx[3], nx[3]},
		{"Empty list", LinkedList{}, nx[0], nx[0]},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.list.PushFront(tt.args)
			if tt.expect != tt.list.head {
				t.Error("Expected:", tt.expect, "Got:", tt.list.head)
			}
		})

	}
}

func TestLinkedList_PushBack(t *testing.T) {
	ll, nx := testData()

	tests := []struct {
		name   string
		list   LinkedList
		args   []*Node
		expect *Node
	}{
		{"Test tail", ll, []*Node{nx[1], nx[2]}, nx[2]},
		{"Test tail 2", ll, []*Node{nx[4], nx[3]}, nx[3]},
		{"Empty list", LinkedList{}, []*Node{nx[0]}, nx[0]},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.list
			for _, v := range tt.args {
				l.PushBack(v)
			}
			if tt.expect != l.tail {
				t.Error("Expected:", tt.expect, "Got:", l.tail)
			}
		})
	}
}

func TestLinkedList_Delete(t *testing.T) {
	ll, _ := testData()

	type args struct {
		key int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "non existent key", args: args{key: 1}},
		{name: "head key", args: args{key: 10}},
		{name: "tail key", args: args{key: 50}},
		{name: "key 30", args: args{key: 30}},
		{name: "key 20", args: args{key: 20}},
		{name: "key 40", args: args{key: 40}},
		{name: "empty list", args: args{key: 40}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll.Delete(tt.args.key)
		})
	}
}

func TestLinkedList_Front(t *testing.T) {
	ll, nx := testData()

	for k, v := range nx {
		t.Run(fmt.Sprintf("Front key: %v, value: %v", k, v.Data), func(t *testing.T) {
			got, err := ll.Front()
			if err != nil {
				t.Errorf("LinkedList.Front() error = %v, wantErr %v", err, nil)
				return
			}
			if got != v.Data {
				t.Errorf("LinkedList.Front() = %v, want %v", got, v.Data)
			}
		})
		ll.Delete(v.Data)
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
	ll, nx := testData()

	for i := 0; i < len(nx); i++ {
		v := ll.tail.Data
		fmt.Println("start list")
		ll.Display()

		t.Run(fmt.Sprintf("Back value: %v", v), func(t *testing.T) {
			got, err := ll.Back()
			if err != nil {
				t.Errorf("LinkedList.Back() error = %v, wantErr %v", err, nil)
				return
			}
			if got != v {
				t.Errorf("LinkedList.Back() = %v, want %v", got, v)
			}
			ll.Delete(got)
			fmt.Println("end list")
			ll.Display()
			fmt.Println("-----")
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
	ll, nx := testData()

	t.Run("Simple reverse", func(t *testing.T) {
		ll.Reverse()
		if ll.head != nx[len(nx)-1] || ll.tail != nx[0] {
			t.Errorf("LinkedList.Reverse() values were not reversed")
		}
	})
}

// testData is a helper function to create a base LinkedList for running tests against.
func testData() (LinkedList, []*Node) {
	ll := LinkedList{}
	nx := []*Node{
		{Data: 10},
		{Data: 20},
		{Data: 30},
		{Data: 40},
		{Data: 50},
	}

	for _, v := range nx {
		ll.PushBack(v)
	}

	return ll, nx
}
