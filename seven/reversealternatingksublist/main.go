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
	printListNode(reverseAlternatingKElementsSublist(&a, 2))
}

func printListNode(q *ListNode) {
	for q != nil {
		fmt.Println(q.Val)
		q = q.Next
	}
}

func reverseAlternatingKElementsSublist(head *ListNode, k int) *ListNode {
	if k <= 1 || head == nil {
		return head
	}

	current := head.Next
	var previous *ListNode
	count := 1

	for current != nil {
		count++
		current = current.Next
	}

	current = head
	count = count / k

	for j := 0; j < count; j++ {
		stopLeft := previous
		stopRight := current

		for i := 0; current != nil && i < k; i++ {
			next := current.Next
			current.Next = previous
			previous = current
			current = next
		}

		if stopLeft != nil {
			stopLeft.Next = previous
		} else {
			head = previous
		}

		stopRight.Next = current
		previous = stopRight
	}
	return head
}
