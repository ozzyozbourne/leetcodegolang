package main

import "math"

func main() {

}

func longestOnes(nums []int, k int) int {
	var st, mxf, mxl int
	for i := 0; i < len(nums); i++ {

		if nums[i] == 1 {
			mxf++
		}

		if i-st+1-mxf > k {
			if nums[st] == 1 {
				mxf--
			}
			st++
		}

		mxl = int(math.Max(float64(mxl), float64(i-st+1)))
	}
	return mxl
}
