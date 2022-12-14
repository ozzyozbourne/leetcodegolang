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

	return head
}
