package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a0ne := ListNode{
		Val:  1,
		Next: nil,
	}

	aTwo := ListNode{
		Val:  2,
		Next: nil,
	}

	a0ne.Next = &aTwo
	aTwo.Next = &a0ne

	fmt.Println(hasCycle(&a0ne))
}

func hasCycle(head *ListNode) bool {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			return true
		}
	}
	return false
}
