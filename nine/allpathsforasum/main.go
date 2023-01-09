package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	a := TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val:   13,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}
	r := pathSum(&a, 22)
	fmt.Println(r)
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	return findPathRec(root, targetSum, make([]int, 0), make([][]int, 0))
}

func findPathRec(currentNode *TreeNode, sum int, currentPath []int, allPaths [][]int) [][]int {

	if currentNode == nil {
		return allPaths
	}
	currentPath = append(currentPath, currentNode.Val)

	if currentNode.Val == sum && currentNode.Left == nil && currentNode.Right == nil {
		//copySlice := make([]int, len(currentPath))
		//copy(copySlice, currentPath)
		allPaths = append(allPaths, currentPath)
	} else {
		allPaths = findPathRec(currentNode.Left, sum-currentNode.Val, currentPath, allPaths)
		allPaths = findPathRec(currentNode.Right, sum-currentNode.Val, currentPath, allPaths)
	}

	return allPaths
}
