package main

import "fmt"

func main() {
	var a int32 = 5
	var b int32 = 15
	var s int32 = 7
	var t int32 = 11
	ap := []int32{-2, 2, 1}
	ora := []int32{5, -6}

	countApplesAndOranges(s, t, a, b, ap, ora)

}

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	var contM int32
	var contL int32

	for _, d := range apples {
		if posFinal := a + d; posFinal >= s && posFinal <= t {
			contM++
		}
	}
	for _, o := range oranges {
		if posFinal := b + o; posFinal >= s && posFinal <= t {
			contL++
		}
	}
	fmt.Println(contM)
	fmt.Println(contL)
}

/*
   Complexidade de Tempo: O(m + n)
   Onde 'm' é o número de maçãs e 'n' o número de laranjas.
   Percorremos cada lista de frutas exatamente uma vez.

   Complexidade de Espaço: O(1)
   Usamos apenas dois contadores fixos, independente da
   quantidade de frutas processadas.
*/
