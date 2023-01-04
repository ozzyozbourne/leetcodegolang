package main

import "math"

func main() {

}

func maxSubArray(nums []int) int {
	l := len(nums)
	var lmax int
	gmax := math.MinInt
	for i := 0; i < l; i++ {
		lmax = int(math.Max(float64(nums[i]), float64(nums[i]+lmax)))
		if lmax > gmax {
			gmax = lmax
		}
	}
	return gmax
}
