package main

func main() {

}

func moveZeroes(nums []int) {
	l := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			t := nums[l]
			nums[l] = nums[i]
			nums[i] = t
			l++
		}
	}
}
