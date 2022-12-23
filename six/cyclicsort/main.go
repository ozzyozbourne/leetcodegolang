package main

import "fmt"

func main() {
	a := []int{3, 1, 5, 4, 2}
	sort(a)
}

func sort(a []int) {
	for i, v := range a {
		if a[v-1] != v {
			swap(a, i, v-1)
		}
	}

	for _, v := range a {
		fmt.Println(v)
	}
}

func swap(a []int, i int, j int) {
	temp := a[i]
	a[i] = a[j]
	a[j] = temp
}
