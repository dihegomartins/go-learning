package main

import "fmt"

func main()  {
	ponto := 85

	if ponto >= 90 {
		fmt.Printf("Cliente VIP")
	} else if ponto >= 70 {
		fmt.Printf("Cliente Premium")
	} else {
		fmt.Printf("Cliente Standard")
	}

	if idade := 18; idade >= 18 {
		fmt.Printf("Pode entrar no sistema")
	}
}