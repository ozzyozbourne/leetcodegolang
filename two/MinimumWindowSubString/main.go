package main

import "fmt"

func main() {
	fmt.Println(minWindow("cabwefgewcwaefgcf", "cae"))
}

func minWindow(s string, t string) string {
	fmap, wmap := make(map[int32]int), make(map[int32]int)
	min, st, sb := len(s)+1, 0, 0
	for _, v := range t {
		fmap[v] += 1
	}
	have, need := 0, len(fmap)
	for i, v := range s {
		wmap[v] += 1
		if _, ok := fmap[v]; ok && fmap[v] == wmap[v] {
			have++
		}

		for need == have {
			if min > i-st+1 {
				min = i - st + 1
				sb = st
			}
			wmap[int32(s[st])] -= 1

			if _, ok := fmap[int32(s[st])]; ok && fmap[int32(s[st])] > wmap[int32(s[st])] {
				have--
			}
			st++
		}
	}
	if min > len(s) {
		return ""
	} else {
		return s[sb : sb+min]
	}

}
