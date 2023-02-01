package main

import "fmt"

func main() {
	fmt.Println(backspaceCompare("xywrrmp", "xywrrmu#p"))
}

func backspaceCompare(s string, t string) bool {
	index1, index2 := len(s)-1, len(t)-1
	for index1 >= 0 || index2 >= 0 {
		i1 := getNextValidCharIndex(s, index1)
		i2 := getNextValidCharIndex(t, index2)
		if i1 < 0 && i2 < 0 {
			return true
		}
		if i1 < 0 || i2 < 0 {
			return false
		}
		if s[i1] != t[i2] {
			return false
		}
		index1 = i1 - 1
		index2 = i2 - 1
	}
	return true
}

func getNextValidCharIndex(s string, index int) int {
	var backSpaceCount int
	for index >= 0 {
		if s[index] == 35 {
			backSpaceCount++
		} else if backSpaceCount > 0 {
			backSpaceCount--
		} else {
			break
		}
		index--
	}
	return index
}
