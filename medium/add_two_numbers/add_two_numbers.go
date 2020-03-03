package main

// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order and each of their nodes contain a single digit.
// Add the two numbers and return it as a linked list.

// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	startNode := &ListNode{}
	l3 := startNode
	carry := false
	for l1 != nil || l2 != nil || carry {
		if l1 != nil {
			l3.Val += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			l3.Val += l2.Val
			l2 = l2.Next
		}

		if carry {
			carry = false
			l3.Val++
		}

		if l3.Val > 9 {
			carry = true
			l3.Val -= 10
		}

		if l1 != nil || l2 != nil || carry {
			l3.Next = &ListNode{}
		}

		l3 = l3.Next
	}

	return startNode
}

func addTwoNumbersTwo(l1, l2 *ListNode) *ListNode {
	startNode := &ListNode{}
	l3 := startNode
	carry := 0
	for l1 != nil || l2 != nil {
		var x, y int
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}

		sum := carry + x + y
		carry = sum / 10
		l3.Next = &ListNode{Val: sum % 10}
		l3 = l3.Next
	}

	if carry > 0 {
		l3.Next = &ListNode{Val: carry}
	}

	return startNode.Next
}
