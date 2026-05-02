package main

import "fmt"

func main() {
	var saldo float64 = 150.50
	var valorSaque float64 = 50

	if resto := (saldo - valorSaque); valorSaque > saldo {
		fmt.Println("Erro: Saldo insuficiente!")
	} else if valorSaque <= 0 {
		fmt.Println("Erro: O valor do saque deve ser positivo!")
	} else {
		fmt.Printf("Saque realizado! Saldo restante: R$ %.2f\n", resto)
	}
}
