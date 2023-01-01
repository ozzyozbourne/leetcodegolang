package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	f := head
	s := head

	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
	}

	rv := reverseList(s)
	cRv := rv

	for head != nil && rv != nil {

		if head.Val != rv.Val {
			break
		}
		head = head.Next
		rv = rv.Next
	}

	reverseList(cRv)
	if head == nil || rv == nil {
		return true
	}

	return false
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
