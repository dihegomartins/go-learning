package main

import "fmt"

func main() {
	//var aprovados map[int]string
	// mapas devem ser inicializados
	aprovados := make(map[int]string)
	aprovados[12345678978] = "Maria"
	aprovados[98785454454] = "Pedro"
	aprovados[55445454512] = "Anna"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[55445454512])
	delete(aprovados, 55445454512)
	fmt.Println(aprovados)

}
