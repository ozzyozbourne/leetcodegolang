package main

func main() {

}

func sortedSquares(nums []int) []int {
	a := make([]int, len(nums))
	l := 0
	r := len(a) - 1
	h := r
	for l <= r {
		sl := nums[l] * nums[l]
		sr := nums[r] * nums[r]
		if sl > sr {
			a[h] = sl
			l++
		} else {
			a[h] = sr
			r--
		}
		h--
	}
	return a
}
