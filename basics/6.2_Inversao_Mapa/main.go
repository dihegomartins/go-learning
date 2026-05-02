package main

import "fmt"

func main() {
	usuarios := map[int]string{1: "Dihego", 2: "Korp", 3: "Angular"}
	nomesParaIDs := make(map[string]int)

	for id, nome := range usuarios {
		nomesParaIDs[nome] = id
		fmt.Printf("A chave: %s | Valor: %v\n", nome, id)
	}

	fmt.Printf("Só um teste no novo maps: %v\n", nomesParaIDs["Dihego"])
}
