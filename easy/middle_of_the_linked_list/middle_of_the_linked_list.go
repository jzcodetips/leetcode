package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	n := getLength(head)

	pos := n/2 + 1

	return getNode(head, pos, 1)
}

func getNode(node *ListNode, pos, curr int) *ListNode {
	if curr < pos {
		return getNode(node.Next, pos, curr+1)
	}

	return node
}

func getLength(node *ListNode) int {
	if node.Next != nil {
		return getLength(node.Next) + 1
	}

	return 1
}
