package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange([]int{0}, 0))
	fmt.Println(searchRange([]int{}, 0))
	fmt.Println(searchRange([]int{2, 2}, 2))
	fmt.Println(searchRange([]int{6, 6, 6}, 6))
	fmt.Println(searchRange([]int{1, 6, 6, 6}, 6))
}

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	res[0] = search(nums, target, false)
	if res[0] != -1 {
		res[1] = search(nums, target, true)
	}
	return res
}

func search(nums []int, target int, b bool) int {

	start, end, keyIndex := 0, len(nums)-1, -1

	for start <= end {
		mid := start + (end-start)/2
		if target > nums[mid] {
			start = mid + 1
		} else if target < nums[mid] {
			end = mid - 1
		} else {
			keyIndex = mid
			if b {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return keyIndex
}
