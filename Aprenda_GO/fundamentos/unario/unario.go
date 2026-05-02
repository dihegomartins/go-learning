package main

import "fmt"

func main() {

	x := 1
	y := 2

	// apenas postfix
	x++ // x += 1 ou x = x + 1
	y++
	fmt.Println(x, y)
	y--
	fmt.Println(y)
}
