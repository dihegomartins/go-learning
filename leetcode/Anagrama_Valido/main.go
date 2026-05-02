package main

import "fmt"

func main() {
	s := "anagram"
	t := "nagaram"
	resultado := isAnagram(s, t)
	fmt.Println(resultado)
}

func isAnagram(s string, t string) bool {
	contagem := map[rune]int{}

	if len(s) != len(t) {
		return false
	}

	for _, letra := range s {
		contagem[letra]++
	}

	for _, letra := range t {
		contagem[letra]--
		if contagem[letra] < 0 {
			return false
		}
	}
	return true
}
