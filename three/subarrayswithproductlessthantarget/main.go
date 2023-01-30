package main

import "fmt"

func main() {
	fmt.Println(numSubarrayProductLessThanK([]int{1, 2, 3}, 0))
}

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	var left, result int
	product := 1
	for right := 0; right < len(nums); right++ {
		product *= nums[right]
		for product >= k {
			product /= nums[left]
			left++
		}
		result += right - left + 1
	}
	return result
}
