package main

import "fmt"

func main() {
	fmt.Println(firstMissingPositive([]int{-3, 1, 5, 4, 2}))
}

func firstMissingPositive(a []int) int {
	i := 0
	l := len(a)
	for i < l {
		if a[i] > 0 && a[i] < l && a[i] != a[a[i]-1] {
			temp := a[i]
			a[i] = a[temp-1]
			a[temp-1] = temp
		} else {
			i++
		}
	}
	for i, v := range a {
		if v != i+1 {
			return i + 1
		}
	}
	return l + 1
}
