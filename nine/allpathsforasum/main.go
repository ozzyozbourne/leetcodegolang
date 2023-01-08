package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	a := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}
	r := pathSum(&a, 3)
	fmt.Println(r)
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	allPaths := make([][]int, 0)
	currentPath := make([]int, 0)
	return findPathRec(root, targetSum, currentPath, allPaths)
}

func findPathRec(currentNode *TreeNode, sum int, currentPath []int, allPaths [][]int) [][]int {

	if currentNode == nil {
		return allPaths
	}
	currentPath = append(currentPath, currentNode.Val)

	if currentNode.Val == sum && currentNode.Left == nil && currentNode.Right == nil {
		copySlice := make([]int, len(currentPath))
		copy(copySlice, currentPath)
		allPaths = append(allPaths, copySlice)
	} else {
		allPaths = findPathRec(currentNode.Left, sum-currentNode.Val, currentPath, allPaths)
		allPaths = findPathRec(currentNode.Right, sum-currentNode.Val, currentPath, allPaths)
	}

	return allPaths
}
