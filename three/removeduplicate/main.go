package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{1, 2, 3, 3, 2}))
}

func removeDuplicates(nums []int) int {
	nextNonDuplicate := 1
	for i := 1; i < len(nums); i++ {
		if nums[nextNonDuplicate-1] != nums[i] {
			nums[nextNonDuplicate] = nums[i]
			nextNonDuplicate++
		}
	}
	return nextNonDuplicate
}
