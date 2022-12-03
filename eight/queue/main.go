package main

import (
	"fmt"
)

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

	fmt.Println(traverse(&val))
}

func traverse(root *TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}

	var queue []*TreeNode
	queue = append(queue, root)

	var result [][]int

	for len(queue) > 0 {

		sz := len(queue)
		var level []int

		for i := 0; i < sz; i++ {

			node := queue[0]
			queue = queue[1:]

			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
	}
	return result
}
