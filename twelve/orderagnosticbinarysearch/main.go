package main

func main() {

}

func search(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := start + (end-start)/2
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}
