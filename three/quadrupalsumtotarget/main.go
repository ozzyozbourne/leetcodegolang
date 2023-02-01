package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(fourSum([]int{4, -9, -2, -2, -7, 9, 9, 5, 10, -10, 4, 5, 2, -4, -2, 4, -9, 5}, -13))
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var resFinal [][]int
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			res := searchPairs(nums, target, i, j)
			if res != nil {
				for _, v := range res {
					resFinal = append(resFinal, v)
				}
			}
		}
	}
	return resFinal
}

func searchPairs(nums []int, target int, first int, second int) [][]int {
	var res [][]int
	left, right := second+1, len(nums)-1
	for left < right {
		sum := nums[first] + nums[second] + nums[left] + nums[right]
		if sum == target {
			res = append(res, []int{nums[first], nums[second], nums[left], nums[right]})
			left++
			right--
			for left < right && nums[left] == nums[left-1] {
				left++
			}
			for left < right && nums[right] == nums[right+1] {
				right--
			}
		} else if sum < target {
			left++
		} else {
			right--
		}

	}
	return res
}
