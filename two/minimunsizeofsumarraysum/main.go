package main

import "math"

func main() {

}

func minSubArrayLen(target int, nums []int) int {
	var ws, wsum int
	sum := math.MaxInt
	for i := 0; i < len(nums); i++ {
		wsum += nums[i]
		for wsum >= target {
			sum = int(math.Min(float64(sum), float64(i-ws+1)))
			wsum -= nums[ws]
			ws++
		}
	}
	if sum == math.MinInt {
		return 0
	}
	return sum
}
