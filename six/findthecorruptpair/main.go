package main

import "fmt"

func main() {
	a := corruptPair([]int{3, 1, 2, 5, 2})
	for _, v := range a {
		fmt.Print(v, "\t")
	}
}

func corruptPair(a []int) []int {
	i := 0
	for i < len(a) {
		if a[i] != a[a[i]-1] {
			temp := a[i]
			a[i] = a[temp-1]
			a[temp-1] = temp
		} else {
			i++
		}
	}
	var r []int
	for i, v := range a {
		if v != i+1 {
			r = append(r, v, i+1)
		}
	}
	return r
}
