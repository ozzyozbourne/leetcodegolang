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
	printListNode(rotateLinkedList(&a, 2))
}

func printListNode(q *ListNode) {
	for q != nil {
		fmt.Println(q.Val)
		q = q.Next
	}
}

func rotateLinkedList(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k < 0 {
		return head
	}

	lastNode := head
	count := 1

	for lastNode.Next != nil {
		count++
		lastNode = lastNode.Next
	}

	lastNode.Next = head
	k %= count
	counter := count - k
	breaker := head

	for i := 0; i < counter-1; i++ {
		breaker = breaker.Next
	}

	head = breaker.Next
	breaker.Next = nil

	return head
}
