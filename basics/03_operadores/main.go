package main

import "fmt"

func main() {
	num1 := 10
	num2 := 3

	soma := num1 + num2
	sub := num1 - num2
	mult := num1 * num2
	div := float64(num1) / float64(num2)

	fmt.Printf("A soma e: %v\nA subtração e: %v\nA multiplicação é: %v\nA divisão é: %.2f\n", soma, sub, mult, div)

}