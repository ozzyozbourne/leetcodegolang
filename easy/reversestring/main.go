package main

func main() {

}

func reverseString(s []byte) {
	var i int
	j := len(s) - 1
	for i < j {
		t := s[i]
		s[i] = s[j]
		s[j] = t
		i++
		j--
	}
}
