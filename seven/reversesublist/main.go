package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a := ListNode{
		Val: 10,
		Next: &ListNode{
			Val: 20,
			Next: &ListNode{
				Val: 30,
				Next: &ListNode{
					Val:  40,
					Next: nil,
				},
			},
		},
	}
	printListNode(reverseBetween(&a, 2, 3))
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {

	if left == right {
		return head
	}

	current := head
	var previous *ListNode

	for i := 0; current != nil && i < left-1; i++ {
		previous = current
		current = current.Next
	}

	leftStop := previous
	rightStop := current

	for i := 0; current != nil && i < right-left+1; i++ {
		next := current.Next
		current.Next = previous
		previous = current
		current = next
	}

	if leftStop != nil {
		leftStop.Next = previous
	} else {
		head = previous
	}

	rightStop.Next = current

	return head
}

func printListNode(q *ListNode) {
	for q != nil {
		fmt.Println(q.Val)
		q = q.Next
	}
}
