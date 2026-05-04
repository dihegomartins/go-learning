package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := []int{555, 901, 482, 1771}
	resultado := findNumbers(s)
	fmt.Println(resultado)
}

func findNumbers(nums []int) int {
	countPar := 0
	for _, itemChar := range nums {
		t := strconv.Itoa(itemChar)
		if len(t)%2 == 0 {
			countPar++
		}
	}
	return countPar
}
