package main

import "fmt"

func main() {
	a := findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1})
	for _, v := range a {
		fmt.Print(v, "\t")
	}
}

func findDuplicates(nums []int) []int {
	i := 0
	for i < len(nums) {
		if nums[i] != nums[nums[i]-1] {
			t := nums[i]
			nums[i] = nums[t-1]
			nums[t-1] = t
		} else {
			i++
		}
	}
	var r []int
	for i, v := range nums {
		if v != i+1 {
			r = append(r, v)
		}
	}
	return r
}
