package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func detectCycle(head *ListNode) *ListNode {
	var l int
	s := head
	f := head
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
		if s == f {
			l = calCycleLength(s)
			break
		}
	}
	return findStart(head, l)
}

func calCycleLength(s *ListNode) int {
	c := s
	var l int
	for {
		c = c.Next
		l++
		if s == c {
			break
		}
	}
	return l
}

func findStart(h *ListNode, c int) *ListNode {
	if h == nil || h.Next == nil || c == 0 {
		return nil
	}
	p1 := h
	p2 := h
	for ; c > 0; c-- {
		p1 = p1.Next
	}

	for p1 != p2 {
		p2 = p2.Next
		p1 = p1.Next
	}
	return p1
}
