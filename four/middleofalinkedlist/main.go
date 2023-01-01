package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	printList(middleNode(&a))
}

func printList(a *ListNode) {
	for a != nil {
		fmt.Print(a.Val, "\t")
		a = a.Next
	}
}

func middleNode(head *ListNode) *ListNode {
	f := head
	s := head
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
	}
	return s
}
