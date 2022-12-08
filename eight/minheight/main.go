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

	fmt.Println(minDepth(&val))
}

func minDepth(root *TreeNode) int {

	var result int
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {

		result++
		sz := len(queue)
		for i := 0; i < sz; i++ {
			cr := queue[0]
			queue = queue[1:]

			if cr.Left == nil && cr.Right == nil {
				return result
			}

			if cr.Left != nil {
				queue = append(queue, cr.Left)
			}

			if cr.Right != nil {
				queue = append(queue, cr.Right)
			}
		}

	}
	return result
}
