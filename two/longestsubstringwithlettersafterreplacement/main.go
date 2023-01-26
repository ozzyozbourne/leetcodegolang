package main

import "math"

func main() {

}

func characterReplacement(s string, k int) int {
	var st, mxf, mxl int
	mp := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		mp[s[i]] = mp[s[i]] + 1
		mxf = int(math.Max(float64(mxf), float64(mp[s[i]])))

		if (i - st + 1 - mxf) > k {
			mp[s[st]] = mp[s[st]] - 1
			st++
		}

		mxl = int(math.Max(float64(mxl), float64(i-st+1)))
	}
	return mxl
}
