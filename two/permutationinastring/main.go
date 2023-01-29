package main

import "fmt"

func main() {
	fmt.Println(checkInclusion("ab", "eidboaoo"))
}

func checkInclusion(s1 string, s2 string) bool {
	fMap := make(map[int32]int)
	for _, v := range s1 {
		fMap[v] += 1
	}
	var st, max int
	for i, v := range s2 {
		if _, ok := fMap[v]; ok {
			fMap[v] -= 1
			if fMap[v] == 0 {
				max++
			}
		}

		if max == len(fMap) {
			return true
		}

		if i >= len(s1)-1 {
			if _, ok := fMap[int32(s2[st])]; ok {
				if fMap[int32(s2[st])] == 0 {
					max++
				}
				fMap[int32(s2[st])] += 1
			}
			st++
		}
	}
	return false
}
