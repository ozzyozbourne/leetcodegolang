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

	fmt.Println(levelOrder(&val, 12))
}

func levelOrder(root *TreeNode, key int) *TreeNode {

	if root == nil {
		return nil
	}
	q := []*TreeNode{root}

	for len(q) > 0 {
		cr := q[0]
		q = q[1:]

		if cr.Left != nil {
			q = append(q, cr.Left)
		}

		if cr.Right != nil {
			q = append(q, cr.Right)
		}

		if cr.Val == key {
			break
		}

	}
	return q[0]
}
