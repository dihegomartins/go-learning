package main

import "fmt"

func main() {
	minhaString := "aabbcdeeff"
	resultado := firstUniqChar(minhaString)

	fmt.Println(resultado)
}

func firstUniqChar(s string) int {
	contagem := make(map[rune]int)

	for i := 0; i < len(s); i++ {
		contagem[rune(s[i])]++
	}

	for i := 0; i < len(s); i++ {
		if contagem[rune(s[i])] == 1 {
			return i
		}
	}
	return -1
}
