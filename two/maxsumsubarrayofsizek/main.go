package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxSubArrayOfSizeK([]int{2, 1, 5, 1, 3, 2}, 3))
}
func maxSubArrayOfSizeK(nums []int, k int) int {
	var wsum, sum, ws int
	for i := 0; i < len(nums); i++ {
		wsum += nums[i]
		if i >= k-1 {
			sum = int(math.Max(float64(sum), float64(wsum)))
			wsum -= nums[ws]
			ws++
		}
	}
	return sum
}
