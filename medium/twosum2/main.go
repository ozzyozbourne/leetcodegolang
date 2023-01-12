package main

func main() {

}

func twoSum(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1
	for l < r {
		s := numbers[l] + numbers[r]
		if s > target {
			r--
		} else if s < target {
			l++
		} else {
			return []int{l + 1, r + 1}
		}
	}
	return []int{}
}

func twoSumDiff(numbers []int, target int) []int {
	var l, s int
	r := len(numbers) - 1
	s = numbers[l] + numbers[r]
	for s != target {
		if s > target {
			r--
		} else {
			l++
		}
		s = numbers[l] + numbers[r]
	}
	return []int{l + 1, r + 1}
}
