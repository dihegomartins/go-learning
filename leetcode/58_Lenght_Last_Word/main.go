package main

import (
	"fmt"
	"strings"
)

func main() {
	//minhaString := "Hello World"
	minhaString := "   fly me   to   the moon  aabbccP"
	fmt.Println(lengthOfLastWord(minhaString))
}

func lengthOfLastWord(s string) int {
	palavras := strings.Fields(s)

	if len(palavras) == 0 {
		return 0
	}

	ultimaPalavra := palavras[len(palavras)-1]
	return len(ultimaPalavra)
}
