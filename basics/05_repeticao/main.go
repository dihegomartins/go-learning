package main

import "fmt"

func main() {

	//1. O for clássico (estilo C)
	for i := 0; i < 5; i++ {
		fmt.Println("Contando:", i)
	}

	//2. O for como while
	//No Go, se você tira o início e o incremento, o for vira um while.
	contador := 0
	for contador < 5 {
		fmt.Println(contador)
		contador++
	}

	//3. O range (O que você mais vai usar no seu projeto da Korp!)
	//Esse é essencial para percorrer listas (slices) de produtos ou notas.

	produtos := []string{"Cimento", "Piso", "Argamassa"}
	for indice, nome := range produtos {
		fmt.Printf("Produto %d: %s\n", indice, nome)
	}
}
