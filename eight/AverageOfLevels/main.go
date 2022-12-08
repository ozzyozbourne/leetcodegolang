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

	fmt.Println(averageOfLevels(&val))
}

func averageOfLevels(root *TreeNode) []float64 {

	var r []float64
	if root == nil {
		return r
	}

	q := []*TreeNode{root}
	for len(q) > 0 {
		sz := len(q)
		var cv float64

		for i := 0; i < sz; i++ {
			cr := q[0]
			q = q[1:]

			cv += float64(cr.Val)
			if cr.Left != nil {
				q = append(q, cr.Left)
			}

			if cr.Right != nil {
				q = append(q, cr.Right)
			}
		}
		r = append(r, cv/float64(sz))
	}
	return r
}
