package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func printTree(root *Node) {
	fmt.Println("Printing")
	for root != nil {
		fmt.Println(root.Val)
		root = root.Next
	}
}

func main() {
	root := Node{
		Val: 12,
		Left: &Node{
			Val: 7,
			Left: &Node{
				Val:   9,
				Left:  nil,
				Right: nil,
				Next:  nil,
			},
			Right: nil,
			Next:  nil,
		},
		Right: &Node{
			Val: 1,
			Left: &Node{
				Val:   10,
				Left:  nil,
				Right: nil,
				Next:  nil,
			},
			Right: &Node{
				Val:   5,
				Left:  nil,
				Right: nil,
				Next:  nil,
			},
			Next: nil,
		},
		Next: nil,
	}

	connect(&root)
	printTree(&root)

}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	q := []*Node{root}
	var p *Node
	for len(q) > 0 {
		if p != nil {
			p.Next = q[0]
		}
		p = q[0]
		q = q[1:]
		if p.Left != nil {
			q = append(q, p.Left)
		}
		if p.Right != nil {
			q = append(q, p.Right)
		}
	}
	return root
}
