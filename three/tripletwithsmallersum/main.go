package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(searchTriplets([]int{-1, 0, 2, 3}, 3))
	fmt.Println(searchTriplets([]int{-1, 4, 2, 1, 3}, 5))

}

func searchTriplets(arr []int, target int) int {
	sort.Ints(arr)
	var count int
	for i := 0; i < len(arr)-2; i++ {
		count += searchPair(arr, target-arr[i], i+1)
	}
	return count
}

func searchPair(arr []int, target int, left int) int {
	var count int
	right := len(arr) - 1
	for left < right {
		if arr[left]+arr[right] < target {
			count += right - left
			left++
		} else {
			right--
		}
	}
	return count
}
