package main

import "fmt"

func main() {
	fmt.Println(isHappy(23))
}

func isHappy(n int) bool {
	f := n
	s := n
	for {
		s = findSumSquare(s)
		f = findSumSquare(findSumSquare(f))
		if s == f {
			break
		}
	}
	return s == 1
}

func findSumSquare(n int) int {
	var sum int
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}
