package main

import "fmt"

func main() {
	/*fmt.Println()

	funcionário := make(map[int]string)
	funcionário[0] = "Dihego"
	funcionário[1] = "Lucas"
	funcionário[2] = "Ana"
	delete(funcionário, 0)

	valor, ok := funcionário[1]

	if ok {
		fmt.Println("O valor é:", valor)
	} else {
		fmt.Println("Essa chave não está no mapa!")
	}

	nome, existe := funcionário[99]
	if existe {
		fmt.Println("O valor é:", nome)
	} else {
		fmt.Println("Essa chave não está no mapa!")
	}*/

	//executarPortalKorp()
	executarInventarioKorp()
}

func executarPortalKorp() {
	autorizados := map[int]string{
		101: "Dihego",
		102: "Lucas",
		105: "Ana",
	}

	idTentativo := 1031
	if nome, ok := autorizados[idTentativo]; ok {
		fmt.Printf("Bem-vindo [%s]!\n", nome)
	} else {
		fmt.Printf("Acesso Negado!")
	}
}

func executarInventarioKorp() {
	estoque := map[string]int{"Prego": 100, "Martelo": 10, "Serrote": 5}
	compras := []string{"Prego", "Furadeira", "Serrote"}

	for _, item := range compras {

		if qtdEstoque, ok := estoque[item]; ok {
			fmt.Printf("Temos [%s]! Quantidade: [%v]\n", item, qtdEstoque)
		} else {
			fmt.Printf("Desculpe, não trabalhamos com [%s]\n", item)
		}
	}
}
