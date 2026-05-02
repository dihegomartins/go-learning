package main

import "fmt"

func main() {
	meusDados := []int{1, 2, 3, 4}
	// Você chama a função e guarda o que ela "retornou" em uma variável
	resultado := getConcatenation(meusDados)
	fmt.Println(resultado)
}

func getConcatenation(nums []int) []int {
	nums = append(nums, nums...)
	return nums
}
