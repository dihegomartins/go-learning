package main

import "fmt"

func main() {

	estoque := make(map[string]int)
	estoque["Parafuso"] = 100
	estoque["Prego"] = 0

	quantidade, ok := estoque["Mart"]
	if !ok {
		fmt.Println("Item não encontrado no sistema!")
	} else if quantidade == 0 {
		fmt.Printf("Produto esgotado")
	} else {
		fmt.Println("Quantidade:", quantidade)
	}
}
