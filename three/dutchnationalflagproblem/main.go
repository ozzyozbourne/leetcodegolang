package main

func main() {
	sortColors([]int{1, 2, 0})
}

func sortColors(nums []int) {
	var low, high = 0, len(nums) - 1
	for i := 0; i <= high; {
		if nums[i] == 0 {
			nums[low], nums[i] = nums[i], nums[low]
			low++
			i++
		} else if nums[i] == 1 {
			i++
		} else {
			nums[high], nums[i] = nums[i], nums[high]
			high--
		}
	}
}
