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
						Val: 50,
						Next: &ListNode{
							Val: 60,
							Next: &ListNode{
								Val:  70,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	printListNode(&a)
	fmt.Println()
	printListNode(reverseAlternatingKElementsSublist(&a, 2))
}

func printListNode(q *ListNode) {
	for q != nil {
		fmt.Println(q.Val)
		q = q.Next
	}
}

func reverseAlternatingKElementsSublist(head *ListNode, k int) *ListNode {

	if head == nil || k <= 0 {
		return nil
	}

	current := head
	var previous *ListNode

	for {

		startLeft := previous
		startRight := current

		for i := 0; current != nil && i < k; i++ {
			next := current.Next
			current.Next = previous
			previous = current
			current = next
		}

		if startLeft != nil {
			startLeft.Next = previous
		} else {
			head = previous
		}

		startRight.Next = current

		for i := 0; current != nil && i < k; i++ {
			previous = current
			current = current.Next
		}

		if current == nil {
			break
		}
	}

	return head
}
