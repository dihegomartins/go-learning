package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	if i := numeroAleatorio(); i > 5 { //tb suportado no switch
		fmt.Println("Ganhou!!!")
	} else {
		fmt.Println("Perdeu!!!")
	}
}

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())

	r := rand.New(s)
	return r.Intn(10)
}
