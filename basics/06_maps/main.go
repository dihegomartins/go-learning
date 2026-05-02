package main

import "fmt"

func main() {
	// Sintaxe: make(map[TipoDaChave]TipoDoValor)
	//estoque := make(map[string]int)
	//Adicionar/Atualizar:
	//estoque["Cimento"] = 50
	//Acessar:
	//quantidade := estoque["Cimento"]
	//fmt.Printf("A quantidade é %v", quantidade)
	//Deletar:
	//delete(estoque, "Cimento")

	fruta := make(map[string]float64)
	fruta["Maça"] = 6.50
	fruta["Banana"] = 3.00
	fruta["Pera"] = 5.20
	//precoMaca := fruta["Maça"]
	//fmt.Printf("O preço da maça é: R$ %.2f\n", precoMaca)
	//delete(fruta, "Banana")

	for nome, preco := range fruta {
		if preco > 5.00 {
			fmt.Printf("Fruta: %s | Preço: R$ %.2f\n", nome, preco)
		}
	}

	// Um mapa onde a chave é String (nome) e o valor é Int (idade)
	// Ou já iniciando com valores:
	tabela := map[string]float64{
		"Cimento": 35.50,
		"Cal":     20.00,
	}
	valor, existe := tabela["ProdutoX"]

	if existe {
		fmt.Println("O valor é:", valor)
	} else {
		fmt.Println("Essa chave não está no mapa!")
	}

	for chave, valor := range tabela {
		fmt.Printf("Produto: %s | Preço: %.2f\n", chave, valor)
	}
}
