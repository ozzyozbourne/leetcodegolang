package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(longestSumStringLen("araaci", 2))
	fmt.Println(longestSumStringLen("araaci", 1))
	fmt.Println(longestSumStringLen("cbbebi", 3))

}

func longestSumStringLen(s string, k int) int {
	if s == "" || len(s) < k {
		return -1
	}
	var st, ml int
	mp := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		rr := s[i]
		mp[rr] = mp[rr] + 1
		for len(mp) > k {
			rr = s[st]
			mp[rr] = mp[rr] - 1
			if mp[rr] == 0 {
				delete(mp, rr)
			}
			st++
		}
		ml = int(math.Max(float64(ml), float64(i-st+1)))
	}
	return ml
}
