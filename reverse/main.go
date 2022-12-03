package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	p := append(a[:1], a[:]...)
	fmt.Println(p)
}

func levelOrderBottom(root *TreeNode) [][]int {

	var result [][]int
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)
		var levelVal []int

		for i := 0; i < size; i++ {
			currentNode := queue[0]
			queue = queue[1:]

			if currentNode.Left != nil {
				queue = append(queue, currentNode.Left)
			}

			if currentNode.Right != nil {
				queue = append(queue, currentNode.Right)
			}

			levelVal = append(levelVal, currentNode.Val)
		}

		if len(result) == 0 {
			result = append(result, levelVal)
		} else {
			result = append(result[:1], result[:]...)
			result[0] = levelVal
		}

	}
	return result
}
