package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func sumNumbers(root *TreeNode) int {
	return sumNumbersHelper(root, 0)
}

func sumNumbersHelper(root *TreeNode, s int) int {
	if root == nil {
		return 0
	}
	s = (s * 10) + root.Val
	if root.Left == nil && root.Right == nil {
		return s
	}
	return sumNumbersHelper(root.Left, s) + sumNumbersHelper(root.Right, s)
}
