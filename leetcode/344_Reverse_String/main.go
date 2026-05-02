package main

import "fmt"

func main() {

	//'H','a','n','n','a','h'
	//'h', 'e', 'l', 'l', 'o'
	s := []byte{'H', 'a', 'n', 'n', 'a', 'h'}

	reverseString(s)
}

func reverseString(s []byte) {
	k := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		s[i], s[k] = s[k], s[i]
		k--
	}
	fmt.Println(string(s))
}
