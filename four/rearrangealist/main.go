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
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	printList(reorderList(&a))
}

func reorderList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return nil
	}

	s := head
	f := head

	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
	}

	s = reverseList(s)
	f = head
	for f != nil && s != nil {
		tp := f.Next
		f.Next = s
		f = tp

		tp = s.Next
		s.Next = f
		s = tp
	}

	if f != nil {
		f.Next = nil
	}
	return head
}

func reverseList(h *ListNode) *ListNode {
	var pr *ListNode
	for h != nil {
		tp := h.Next
		h.Next = pr
		pr = h
		h = tp
	}
	return pr
}

func printList(a *ListNode) {
	for a != nil {
		fmt.Print(a.Val, "\t")
		a = a.Next
	}
}
