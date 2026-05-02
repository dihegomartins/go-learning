package main

import "fmt"

func main() {
	array := []int{2, 7, 11, 15}
	target := 17
	has := make(map[int]int)

	for i := 0; i < len(array); i++ {
		complemento := target - array[i]

		indiceComplemento, existe := has[complemento]

		if existe {
			fmt.Printf("Achei! Índices: %d %d\n", indiceComplemento, i)
		}

		has[array[i]] = i
	}
}
