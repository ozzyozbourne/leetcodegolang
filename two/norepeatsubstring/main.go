package main

import "math"

func main() {

}

func lengthOfLongestSubstring(s string) int {
	var st, ml int
	mp := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		if _, ok := mp[s[i]]; ok {
			st = int(math.Max(float64(st), float64(mp[s[i]]+1)))
		}
		mp[s[i]] = i
		ml = int(math.Max(float64(ml), float64(i-st+1)))
	}
	return ml
}
