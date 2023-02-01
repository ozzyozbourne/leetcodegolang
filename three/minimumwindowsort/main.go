package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findUnsortedSubarray([]int{1, 3, 2, 0, -1, 7, 10}))
}

func findUnsortedSubarray(nums []int) int {
	low, high := 0, len(nums)-1
	for low < len(nums)-1 && nums[low] <= nums[low+1] {
		low++
	}
	if low == len(nums)-1 {
		return 0
	}
	for high > 0 && nums[high] >= nums[high-1] {
		high--
	}

	min, max := math.MaxInt, math.MinInt

	for i := low; i <= high; i++ {
		min = int(math.Min(float64(min), float64(nums[i])))
		max = int(math.Max(float64(max), float64(nums[i])))
	}
	for low > 0 && nums[low-1] > min {
		low--
	}
	for high < len(nums)-1 && nums[high+1] < max {
		high++
	}

	return high - low + 1
}
