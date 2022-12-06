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

	fmt.Println(zigzagLevelOrder(&val))
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var result [][]int
	ltr := false
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		sz := len(queue)
		var lv []int

		for i := 0; i < sz; i++ {
			cr := queue[0]
			queue = queue[1:]

			if ltr {
				if len(lv) == 0 {
					lv = append(lv, cr.Val)
				} else {
					lv = append(lv[:1], lv[:]...)
					lv[0] = cr.Val
				}
			} else {
				lv = append(lv, cr.Val)
			}

			if cr.Left != nil {
				queue = append(queue, cr.Left)
			}
			if cr.Right != nil {
				queue = append(queue, cr.Right)
			}

		}

		ltr = !ltr
		result = append(result, lv)
	}

	return result
}
