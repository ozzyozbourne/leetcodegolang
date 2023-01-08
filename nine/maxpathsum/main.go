package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	a := TreeNode{
		Val:   -3,
		Left:  nil,
		Right: nil,
	}
	m, nm := maxPathSumHelper(&a)
	fmt.Println(m, "\t", nm)
}

func maxPathSumHelper(root *TreeNode) (float64, float64) {

	if root == nil {
		return float64(math.MinInt), 0
	}

	nvf := float64(root.Val)
	ml, ls := maxPathSumHelper(root.Left)
	mr, rs := maxPathSumHelper(root.Right)

	m := math.Max(ml, mr)
	m = math.Max(m, nvf+ls+rs)

	return m, nvf + math.Max(ls, rs)
}
