package main

import (
	"math"
	"sort"
)

func main() {

}

func threeSumClosest(nums []int, target int) int {
	result := nums[0] + nums[1] + nums[2]
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1
		for left < right {
			current_result := nums[i] + nums[left] + nums[right]
			if current_result > target {
				right--
			} else {
				left++
			}
			if math.Abs(float64(current_result-target)) < math.Abs(float64(result-target)) {
				result = current_result
			}
		}
	}
	return result
}
