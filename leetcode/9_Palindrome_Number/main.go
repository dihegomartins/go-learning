package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := -12121

	fmt.Println(isPalindrome(num))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	numeroString := strconv.Itoa(x)
	ini := 0
	fim := len(numeroString) - 1

	for ini < fim {
		if numeroString[ini] != numeroString[fim] {
			return false
		}
		ini++
		fim--
	}
	return true
}
