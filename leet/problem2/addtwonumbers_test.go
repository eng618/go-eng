package problem2_test

import (
	"reflect"
	"testing"

	"github.com/eng618/go-eng/leet/problem2"
)

func TestAddTwoNumbers(t *testing.T) {
	e1_1 := &problem2.ListNode{Val: 2, Next: &problem2.ListNode{Val: 4, Next: &problem2.ListNode{Val: 3}}}
	e1_2 := &problem2.ListNode{Val: 5, Next: &problem2.ListNode{Val: 6, Next: &problem2.ListNode{Val: 4}}}
	a1 := &problem2.ListNode{Val: 7, Next: &problem2.ListNode{Val: 0, Next: &problem2.ListNode{Val: 8}}}

	type args struct {
		l1 *problem2.ListNode
		l2 *problem2.ListNode
	}
	tests := []struct {
		name string
		args args
		want *problem2.ListNode
	}{
		{name: "Example 1", args: args{l1: e1_1, l2: e1_2}, want: a1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := problem2.AddTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
