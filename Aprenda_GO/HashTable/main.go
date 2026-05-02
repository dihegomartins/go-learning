package main

import "fmt"

func main() {
	counts := make(map[string]int)

	// Adicionando ou atualizando
	counts["apple"] = 5

	// Verificando se uma chave existe (O famoso "comma ok" idiom)

	_, existe := counts["banana"]
	if !existe {
		fmt.Println("Não tem banana")
	}

}
