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

	fmt.Println(rightSideView(&val))
}

func rightSideView(root *TreeNode) []int {
	var r []int
	if root == nil {
		return r
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		s := len(q)
		var lv int
		for i := 0; i < s; i++ {
			cr := q[0]
			q = q[1:]
			if i == s-1 {
				lv = cr.Val
			}
			if cr.Left != nil {
				q = append(q, cr.Left)
			}
			if cr.Right != nil {
				q = append(q, cr.Right)
			}
		}
		r = append(r, lv)
	}
	return r
}
