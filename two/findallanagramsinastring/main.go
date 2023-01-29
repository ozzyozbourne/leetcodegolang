package main

func main() {

}

func findAnagrams(s string, p string) []int {
	var res []int
	fmap := make(map[int32]int)
	for _, v := range p {
		fmap[v] += +1
	}
	var st, m int
	for i, v := range s {
		if _, ok := fmap[v]; ok {
			fmap[v] -= 1
			if fmap[v] == 0 {
				m++
			}
		}
		if m == len(fmap) {
			res = append(res, st)
		}
		if i >= len(p)-1 {
			if _, ok := fmap[int32(s[st])]; ok {
				if fmap[int32(s[st])] == 0 {
					m--
				}
				fmap[int32(s[st])] += 1
			}
			st++
		}
	}
	return res
}
