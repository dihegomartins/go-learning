/*
	O que é um Slice?

Diferente de um Array (que tem tamanho fixo e você quase não vai usar), o Slice é uma lista dinâmica.
Ele cresce e diminui conforme você precisa.

	Operações essenciais que você precisa treinar:

Declaração: numeros := []int{10, 20, 30}

Adicionar itens (Append): numeros = append(numeros, 40)

Fatiar (Slicing): parte := numeros[1:3] (pega do índice 1 até o 2).

Tamanho: len(numeros)

numeros := []int{10, 20, 30}

	numeros = append(numeros, 33)
	parte := numeros[1:3]

	fmt.Printf("Ola %v", parte)
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	//fmt.Printf("--- Executando Exercício 06 ---\n")
	//executarCarrinho()

	//fmt.Printf("--- Executando Exercício 07 ---\n")
	//executarInversor()

	//fmt.Printf("--- Executando Exercício 08 ---\n")
	//executarFiltroPrecos()

	//fmt.Printf("--- Executando Exercício 09 ---\n")
	//executarAjustePrecos()

	//fmt.Printf("--- Executando Exercício 10 ---\n")
	//executarLimpezaEstoque()

	//fmt.Printf("--- Executando Exercício 11 ---\n")
	//executarAnalisePrecos()

	//fmt.Printf("--- Executando Exercício 12 ---\n")
	//executarUnificacaoEstoque()

	//fmt.Printf("--- Executando Exercício 13 ---\n")
	//executarBuscaProduto()

	//fmt.Printf("--- Executando Exercício 14 ---\n")
	//executarSomaPares()

	//executarLimpezaNomes()

	executarSistemaNotas()
}

func getConcatenation(nums []int) []int {
	nums = append(nums, nums...)
	return nums
}

func executarSistemaNotas() {
	notas := []float64{8.5, -2.0, 5.0, 10.0, -1.5, 7.0, 3.5}
	notasValidas := []float64{}
	sum := 0.0
	for i := 0; i < len(notas); i++ {
		if notas[i] > 0 {
			if notas[i]+1.0 > 10.0 {
				notasValidas = append(notasValidas, 10.0)
				sum += 10.0
			} else {
				sum += (notas[i] + 1.0)
				notasValidas = append(notasValidas, notas[i]+1.0)
			}
		}
	}
	fmt.Println(notasValidas)
	fmt.Printf("A média das notas é: %.2f", sum/float64(len(notasValidas)))
}

func executarLimpezaNomes() {
	nome := []string{" JoAo", "Maria ", " pedro ", "ANA"}

	for i := 0; i < len(nome); i++ {
		palavra := strings.ToLower(nome[i])
		nome[i] = strings.TrimSpace(palavra)
	}

	for _, nomes := range nome {
		fmt.Printf("%v ", nomes)
	}
}

func executarSomaPares() {
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	pares := []int{}
	soma := 0
	for i := 0; i < len(numeros); i++ {
		if numeros[i]%2 == 0 {
			pares = append(pares, numeros[i])
			soma += numeros[i]
		}
	}
	fmt.Printf("A soma total e: %v", soma)
}

func executarBuscaProduto() {
	estoque := []string{"Prego", "Parafuso", "Martelo", "Alicate"}
	alvo := "Martelo"
	encontrado := false
	for i := 0; i < len(estoque); i++ {
		if alvo == estoque[i] {
			fmt.Printf("Produto [%s] encontrado no índice [%v]!", estoque[i], i)
			encontrado = true
			break
		}
	}
	if encontrado == false {
		fmt.Printf("Produto não cadastrado")
	}
}

func executarUnificacaoEstoque() {
	A := []string{"Cimento", "Cal"}
	B := []string{"Areia", "Brita", "Tijolo"}
	combined := append(A, B...)
	fmt.Println(combined)
}

func executarCarrinho() {
	carrinho := []string{}
	carrinho = append(carrinho, "Arroz", "Feijão", "Azeite")
	fmt.Printf("O tamanho do carrinho é: %v\n", len(carrinho))

	favoritos := carrinho[0:2]
	fmt.Printf("Os favoritos: %s\n", favoritos)
}

func executarInversor() {
	lista := []int{1, 2, 3, 4, 5}
	invertido := []int{}

	for i := len(lista) - 1; i >= 0; i-- {
		invertido = append(invertido, lista[i])
	}

	for _, valor := range invertido {
		fmt.Printf("%v ", valor)

	}
}

func executarFiltroPrecos() {
	preco := []float64{10.50, 100.0, 5.25, 250.0, 8.0, 15.0}
	precosAltos := []float64{}

	for _, valor := range preco {
		if valor > 50 {
			precosAltos = append(precosAltos, valor)
		}
	}

	for _, index := range precosAltos {
		fmt.Printf("%.1f ", index)
	}
	fmt.Printf("\nA quantidade de valores adicionando fora %v ", len(precosAltos))

}

func executarAjustePrecos() {
	precos := []float64{10.50, 100.0, 5.25, 250.0, 8.0, 15.0}

	for i := 0; i < len(precos); i++ {
		if precos[i] < 20.0 {
			ajuste := precos[i] * 1.10
			precos[i] = ajuste
		}
	}

	for _, valor := range precos {
		fmt.Printf("%.1f ", valor)
	}
}

func executarLimpezaEstoque() {
	precos := []float64{10.50, 0.0, 5.25, 0.0, 8.0, 15.0}
	estoqueLimpo := []float64{}

	for _, numero := range precos {
		if numero != 0.0 {
			estoqueLimpo = append(estoqueLimpo, numero)
		}
	}

	fmt.Printf("O tamanho do novo slice e: %v\n", len(estoqueLimpo))
	for _, valor := range estoqueLimpo {
		fmt.Printf("%.1f ", valor)
	}

}

func executarAnalisePrecos() {
	precos := []float64{12.50, 45.90, 5.50, 120.0, 8.0, 15.60}
	maior := precos[0]
	menor := precos[0]

	for _, valor := range precos[1:] {
		if valor > maior {
			maior = valor
		} else if valor < menor {
			menor = valor
		}
	}
	fmt.Printf("O maior preço e: R$ %.1f\nO menor preço e: R$ %.1f", maior, menor)
}
