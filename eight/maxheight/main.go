package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	val := TreeNode{
		Val: 12,
		Left: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   10,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
		},
	}

	fmt.Println(maxDepth(&val))
}
func maxDepth(root *TreeNode) int {
	var r int
	if root == nil {
		return r
	}

	q := []*TreeNode{root}

	for len(q) > 0 {
		s := len(q)

		for i := 0; i < s; i++ {
			cr := q[0]
			q = q[1:]

			if cr.Left == nil && cr.Right == nil {
				continue
			}
			if cr.Left != nil {
				q = append(q, cr.Left)
			}
			if cr.Right != nil {
				q = append(q, cr.Right)
			}
		}
		r++
	}
	return r
}
