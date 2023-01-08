package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}
	fmt.Println(binaryTreePathsHelper(&root, "", make([]string, 0)))
}

func binaryTreePathsHelper(root *TreeNode, p string, ps []string) []string {
	if root == nil {
		return ps
	}

	p = p + strconv.Itoa(root.Val) + "->"

	if root.Left == nil && root.Right == nil {
		p = p[:len(p)-2]
		ps = append(ps, p)
	} else {
		ps = binaryTreePathsHelper(root.Left, p, ps)
		ps = binaryTreePathsHelper(root.Right, p, ps)
	}
	return ps
}
