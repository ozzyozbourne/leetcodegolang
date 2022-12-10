package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func main() {

}

func connect(root *Node) *Node {

	if root == nil {
		return nil
	}

	q := []*Node{root}

	for len(q) > 0 {

		s := len(q)
		var p *Node

		for i := 0; i < s; i++ {

			cr := q[0]
			q = q[1:]

			if p != nil {
				p.Next = cr
			}

			p = cr

			if p.Left != nil {
				q = append(q, p.Left)
			}

			if p.Right != nil {
				q = append(q, p.Right)
			}
		}
	}
	return root
}
