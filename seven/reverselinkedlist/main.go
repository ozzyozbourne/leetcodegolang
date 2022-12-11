package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var previous *ListNode

	for head != nil {
		next := head.Next
		head.Next = previous
		previous = head
		head = next
	}
	return previous
}
