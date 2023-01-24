package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	mp := make(map[int]string)
	s, t := mp[75]
	fmt.Println(s)
	fmt.Println(t)
}
