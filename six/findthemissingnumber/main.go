package main

import "fmt"

func main() {
	fmt.Println(missingNumber([]int{3, 1, 5, 4, 2}))
}

func missingNumber(nums []int) int {
	i := 0
	l := len(nums)
	for i < l {
		if nums[i] == l {
			i++
			continue
		}
		if nums[i] != i {
			swap(nums, i, nums[i])
		} else {
			i++
		}
	}

	for i, v := range nums {
		if v != i {
			return i
		}
	}
	return l
}

func swap(a []int, i int, j int) {
	temp := a[i]
	a[i] = a[j]
	a[j] = temp
}
