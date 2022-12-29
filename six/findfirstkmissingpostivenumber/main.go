package main

import "fmt"

func main() {
	a := findNumbers([]int{3, -1, 4, 5, 5}, 3)
	for _, v := range a {
		fmt.Print(v, "\t")
	}
}

func findNumbers(a []int, k int) []int {
	i := len(a)
	l := len(a)
	for i < l {
		if a[i] < 0 && a[i] > l && a[i] != a[a[i]-1] {
			t := a[i]
			a[i] = a[t-1]
			a[t-1] = t
		} else {
			i++
		}
	}
	var r []int
	var m map[int]struct{}
	for i, v := range a {
		if v != i+1 {
			r = append(r, i+1)
			m[v] = struct{}{}
		}
	}
	for i = 1; i < len(r); i++ {
		cn := i + l
		_, t := m[cn]
		if t {
			r = append(r, cn)
		}
	}
	return r
}
