package main

import "fmt"

func main() {
	a := []int{3, 1, 5, 4, 2}
	sort(a)
	for _, v := range a {
		fmt.Println(v)
	}
}

func sort(a []int) {
	i := 0
	l := len(a)
	for i < l {
		if a[i] == l {
			i++
			continue
		}
		if a[i] != i {
			swap(a, i, a[i])
		} else {
			i++
		}
	}
}

func swap(a []int, i int, j int) {
	temp := a[i]
	a[i] = a[j]
	a[j] = temp
}
