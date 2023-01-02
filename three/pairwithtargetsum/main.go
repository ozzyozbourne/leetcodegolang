package main

import "fmt"

func main() {
	for _, v := range twoSum([]int{1, 2}, 4) {
		fmt.Print(v, ",\t")
	}
}

func twoSum(nums []int, target int) []int {
	var r []int
	m := make(map[int]int)
	for i, v := range nums {
		c := target - v
		_, ok := m[c]
		if ok {
			r = append(r, i, m[c])
			return r
		} else {
			m[v] = i
		}
	}
	return r
}
