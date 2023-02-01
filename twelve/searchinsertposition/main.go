package main

func main() {

}
func searchInsert(nums []int, target int) int {
	if target > nums[len(nums)-1] {
		return len(nums)
	}
	start, end := 0, len(nums)-1
	for start <= end {
		mid := start + (end-start)/2
		if target > nums[mid] {
			start = mid + 1
		} else if target < nums[mid] {
			end = mid - 1
		} else {
			return mid
		}
	}
	return start
}
