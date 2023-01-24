package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(totalFruit([]int{1, 2, 3, 1, 3}))
	fmt.Println(totalFruit([]int{1, 2, 3, 2, 2, 3}))
}

func totalFruit(fruits []int) int {
	var st, mx int
	mp := make(map[int]int)
	for i := 0; i < len(fruits); i++ {
		cr := fruits[i]
		mp[cr] = mp[cr] + 1
		for len(mp) > 2 {
			cr = fruits[st]
			mp[cr] = mp[cr] - 1
			if mp[cr] == 0 {
				delete(mp, cr)
			}
			st++
		}
		mx = int(math.Max(float64(mx), float64(i-st+1)))
	}
	return mx
}
