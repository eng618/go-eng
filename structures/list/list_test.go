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

func TestLinkedList_PushBack(t *testing.T) {
	node1 := &Node{Data: 10}
	node2 := &Node{Data: 20}
	node3 := &Node{Data: 40}
	node4 := &Node{Data: 50}
	ll := LinkedList{}

	tests := []struct {
		name   string
		list   LinkedList
		args   []*Node
		expect *Node
	}{
		{"Test tail", ll, []*Node{node1, node2}, node2},
		{"Test tail 2", ll, []*Node{node4, node3}, node3},
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
	ll := testData()

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
		{name: "key 1", args: args{key: 30}},
		{name: "key 1", args: args{key: 20}},
		{name: "key 1", args: args{key: 40}},
		{name: "empty list", args: args{key: 40}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll.Delete(tt.args.key)
		})
	}
}

// testData is a helper function to create a base LinkedList for running tests against.
func testData() LinkedList {
	ll := LinkedList{}

	node1 := &Node{Data: 10}
	node2 := &Node{Data: 20}
	node3 := &Node{Data: 30}
	node4 := &Node{Data: 40}
	node5 := &Node{Data: 50}

	ll.PushBack(node1)
	ll.PushBack(node2)
	ll.PushBack(node3)
	ll.PushBack(node4)
	ll.PushBack(node5)

	return ll
}
