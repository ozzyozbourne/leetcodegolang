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
					Val: 40,
					Next: &ListNode{
						Val:  50,
						Next: nil,
					},
				},
			},
		},
	}
	printListNode(reverseKGroup(&a, 2))
}

func printListNode(q *ListNode) {
	for q != nil {
		fmt.Println(q.Val)
		q = q.Next
	}
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 || head == nil {
		return head
	}

	current := head
	var previous *ListNode

	for {
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

		if current == nil {
			break
		}
		previous = stopRight
	}
	return head
}
