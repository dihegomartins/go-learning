package main

import "fmt"

func main() {
	// Slice de saldos (simulando entrada do LeetCode)
	saldos := []int{10, 25, -5, 40, 0, -2, 15}
	total := 0

	for _, saldo := range saldos {
		if saldo > 0 {
			total += saldo
		}
	}

	fmt.Println("O saldo total positivo é:", total)
}
