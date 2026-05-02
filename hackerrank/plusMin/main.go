package main

import "fmt"

func main() {
	arr := []int32{1, 1, 0, -1, -1}
	plusMinus(arr)
}

func plusMinus(arr []int32) {
	var pos int32
	var neg int32
	var zeros int32
	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			pos++
		} else if arr[i] < 0 {
			neg++
		} else {
			zeros++
		}
	}

	fmt.Printf("%.6f\n%.6f\n%.6f", float64(pos)/float64(len(arr)), float64(neg)/float64(len(arr)), float64(zeros)/float64(len(arr)))
}
