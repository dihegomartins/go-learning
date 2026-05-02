package main

import "fmt"

func main() {
	ar := []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}
	resultado := aVeryBigSum(ar)
	fmt.Println(resultado)
}

func aVeryBigSum(ar []int64) int64 {
	var sum int64
	for _, valor := range ar {
		sum += valor
	}
	return sum
}
