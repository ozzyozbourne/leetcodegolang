package main

import "fmt"

func main() {

	st := make([]string, 4, 5)
	st[0] = "qwd"
	st[1] = "sfdv"
	st[2] = "sd"
	st[3] = "weg"

	sd := append(st, "kdsmc")
	st[0] = "skdnv"

	fmt.Println(sd)
	fmt.Println(st)

}
