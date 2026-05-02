package main

import (
	"fmt"
	"strings"
)

func main() {
	var num int32
	num = 4
	staircase(num)
}

func staircase(n int32) {
	for i := 1; i <= int(n); i++ {
		fmt.Printf("% *s\n", n, strings.Repeat("#", i))
	}
}
