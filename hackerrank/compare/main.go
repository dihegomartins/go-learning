package main

import "fmt"

func main() {
	a := []int32{17, 28, 30}
	b := []int32{99, 16, 8}
	resultado := compareTriplets(a, b)
	fmt.Println(resultado)
}

func compareTriplets(a []int32, b []int32) []int32 {
	var pontA int32
	var pontB int32

	for i := 0; i < 3; i++ {
		if a[i] > b[i] {
			pontA += 1
		} else if a[i] < b[i] {
			pontB += 1
		}
	}

	return []int32{pontA, pontB}
}
