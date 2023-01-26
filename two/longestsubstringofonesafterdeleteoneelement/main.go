package main

import "math"

func main() {

}

func longestSubarray(nums []int) int {
	var st, mxz, mx int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			mxz++
		}
		for mxz > 1 {
			if nums[st] == 0 {
				mxz--
			}
			st++
		}
		mx = int(math.Max(float64(mx), float64(i-st)))
	}
	return mx
}
