package main

import (
	"fmt"
	"slices"
)

func main() {
	var ar []int32
	ar = append(ar, 3, 2, 3, 3)
	resultado := birthdayCakeCandles(ar)
	fmt.Println(resultado)
}

func birthdayCakeCandles(candles []int32) int32 {
	var cont int32
	maximo := slices.Max(candles)

	for i := 0; i < len(candles); i++ {
		if candles[i] == maximo {
			cont++
		}
	}
	return cont
}
