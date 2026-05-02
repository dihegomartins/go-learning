package main

import (
	"fmt"
	"strconv"
)

func main() {
	testArray := []byte{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'b', 'b'}

	resultLen := compress(testArray)

	fmt.Printf("Novo comprimento: %d\n", resultLen)
	fmt.Printf("Array resultante: %s\n", string(testArray[:resultLen]))
}

func compress(chars []byte) int {
	pont1 := 0
	contador := 1

	if len(chars) == 1 {
		return 1
	}

	for i := 1; i <= len(chars); i++ {
		if i < len(chars) && chars[i] == chars[i-1] {
			contador++
		} else {
			chars[pont1] = chars[i-1]
			pont1++

			if contador > 1 {
				s := strconv.Itoa(contador)
				for j := 0; j < len(s); j++ {
					chars[pont1] = s[j]
					pont1++
				}
			}
			contador = 1
		}
	}
	return pont1
}
