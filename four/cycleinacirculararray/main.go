package main

func main() {

}

func circularArrayLoop(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		fw := nums[i] > 0
		var s int
		var f int
		for {
			s = findNextIndex(nums, fw, i)
			f = findNextIndex(nums, fw, i)
			if f != -1 {
				f = findNextIndex(nums, fw, i)

			}

			if s == -1 && f == -1 {
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
		return -1
	}

	return ci
}
