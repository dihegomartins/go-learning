package main

import "fmt"

func main() {
	//x1=0, v1=3, x2=4, v2=2
	var x1 int32 = 0
	var v1 int32 = 3
	var x2 int32 = 4
	var v2 int32 = 2
	resultado := kangaroo(x1, v1, x2, v2)
	fmt.Println(resultado)
}

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	if v1 <= v2 {
		return "NO"
	}
	if (x2-x1)%(v1-v2) == 0 {
		return "YES"
	}
	return "NO"
}

/*
   Complexidade de Tempo: O(1)
   Não há loops. A resposta é calculada instantaneamente com
   operações aritméticas simples, independente da distância.

   Complexidade de Espaço: O(1)
   Não alocamos memória extra; usamos apenas as variáveis de entrada.
*/
