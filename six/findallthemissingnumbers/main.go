package main

import "fmt"

func main() {
	a := findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1})
	for _, v := range a {
		fmt.Print(v, "\t")
	}
}

func findDisappearedNumbers(nums []int) []int {
	i := 0
	var r []int
	for i < len(nums) {
		if nums[i] != nums[nums[i]-1] {
			temp := nums[i]
			nums[i] = nums[temp-1]
			nums[temp-1] = temp
		} else {
			i++
		}
	}
	for i, v := range nums {
		if v != i+1 {
			r = append(r, i+1)
		}
	}
	return r
}
