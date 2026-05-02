package main

import "fmt"

func main() {
	s := []int32{1, 2, 1, 3, 2}
	var d int32 = 3
	var m int32 = 2

	res := birthday(s, d, m)
	fmt.Printf("Resultado esperado: 2 | Resultado obtido: %d\n", res)
}

func birthday(s []int32, d int32, m int32) int32 {
	var contador int32 = 0
	var somaAtual int32 = 0

	if int32(len(s)) < m {
		return 0
	}

	for i := int32(0); i < m; i++ {
		somaAtual += s[i]
	}

	if somaAtual == d {
		contador++
	}

	for i := m; i < int32(len(s)); i++ {
		somaAtual += s[i]

		somaAtual -= s[i-m]

		if somaAtual == d {
			contador++
		}
	}

	return contador
}

/*
   Análise de Complexidade:

   Tempo: O(n) - Percorremos o array uma única vez usando Janela Deslizante.
   Espaço: O(1) - Utilizamos apenas variáveis fixas, independente do tamanho da entrada.
*/
