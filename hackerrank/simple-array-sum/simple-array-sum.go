package main

import "fmt"

func main() {
	ar := []int32{1, 2, 3, 4}

	resultado := simpleArraySum(ar)
	fmt.Println(resultado)
}

func simpleArraySum(ar []int32) int32 {
	var sum int32 = 0
	for _, valor := range ar {
		sum += valor
	}
	return sum
}
