package main

import "fmt"

func main() {
	fmt.Println(circularArrayLoop([]int{2, -1, 1, 2, 2}))
}

func circularArrayLoop(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		fw := nums[i] >= 0
		s := i
		f := i
		for {
			s = findNextIndex(nums, fw, s)
			f = findNextIndex(nums, fw, f)
			if f != -1 {
				f = findNextIndex(nums, fw, f)
			}
			if s == -1 || f == -1 || s == f {
				break
			}
		}
		if s != -1 && f == s {
			return true
		}
	}
	return false
}

func findNextIndex(a []int, f bool, ci int) int {
	if (a[ci] >= 0) != f {
		return -1
	}

	ni := (ci + a[ci]) % len(a)
	if ni < 0 {
		ni += len(a)
	}
	if ni == ci {
		ni = -1
	}

	return ni
}
