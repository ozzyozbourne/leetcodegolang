package main

import "fmt"

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
}

func findDuplicate(nums []int) int {
	i := 0
	for i < len(nums) {
		if nums[i] != i+1 {
			if nums[i] != nums[nums[i]-1] {
				temp := nums[i]
				nums[i] = nums[temp-1]
				nums[temp-1] = temp
			} else {
				return nums[i]
			}
		} else {
			i++
		}
	}
	return -1
}
