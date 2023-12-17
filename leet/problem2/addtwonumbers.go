package problem2

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// AddTwoNumbers is a way to add two numbers that are represented as a linked list in revers order
// https://leetcode.com/problems/add-two-numbers/description/
//
// You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
func AddTwoNumbers(l1, l2 *ListNode) *ListNode {
	answer := new(ListNode)

	// Placeholders
	currNode := answer
	var carry int

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		currNode.Next = &ListNode{
			Val:  sum % 10,
			Next: nil,
		}

		carry = sum / 10
		currNode = currNode.Next
	}

	return answer.Next
}
