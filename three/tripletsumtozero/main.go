package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func threeSum(nums []int) [][]int {
	var res2d [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		res1d := searchPair(nums, -nums[i], i+1)
		if res1d != nil {
			res2d = append(res2d, res1d)
		}
	}
	return res2d
}

func searchPair(arr []int, target int, left int) []int {
	var res []int
	right := len(arr) - 1
	for left < right {
		sum := arr[left] + arr[right]
		if sum == target {
			res = append(res, -target, arr[left], arr[right])
			left++
			right--
			for left < right && arr[left] == arr[left-1] {
				left++
			}
			for left < right && arr[right] == arr[right-1] {
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
