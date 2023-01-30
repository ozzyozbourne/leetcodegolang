package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-2, 0, 1, 1, 2}))
}

func threeSum(nums []int) [][]int {
	var resFinal [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		res := searchPair(nums, -nums[i], i+1)
		if res != nil {
			for _, v := range res {
				resFinal = append(resFinal, v)
			}
		}
	}
	return resFinal
}

func searchPair(arr []int, target int, left int) [][]int {
	var res [][]int
	right := len(arr) - 1
	for left < right {
		sum := arr[left] + arr[right]
		if sum == target {
			res = append(res, []int{-target, arr[left], arr[right]})
			left++
			right--
			for left < right && arr[left] == arr[left-1] {
				left++
			}
			for left < right && arr[right] == arr[right+1] {
				right--
			}

		} else if target > sum {
			left++
		} else {
			right--
		}
	}
	return res
}
