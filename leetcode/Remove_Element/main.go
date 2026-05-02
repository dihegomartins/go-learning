package main

import "fmt"

func main() {
	meusDados := []int{1, 2, 3, 4}
	valor := 1
	// Você chama a função e guarda o que ela "retornou" em uma variável
	resultado := removeElement(meusDados, valor)
	fmt.Println(resultado)
}

func removeElement(nums []int, val int) int {
	k := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}
