package main

import (
	"fmt"
	"slices"
)

func main() {
	arr := []int32{1, 2, 3, 4, 5}

	miniMaxSum(arr)
}

func miniMaxSum(arr []int32) {
	var total int64
	for i := 0; i < len(arr); i++ {
		total += int64(arr[i])
	}
	fmt.Printf("%v %v", total-int64(slices.Max(arr)), total-int64(slices.Min(arr)))
}
