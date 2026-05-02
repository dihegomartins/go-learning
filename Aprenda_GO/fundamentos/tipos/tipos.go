package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(320000))

	// sem sinal (só positivos)... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal.. int8 int16 int32 int64

	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' //representa um mapeamento da tabela Unicode (int32)
	fmt.Println("O tipo de i2 é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// números reias (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))

	// boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))

	// string
	s1 := "Olá meu nome é Dihêgo"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	// string com multiplas linhas
	s2 := ` Olá	
		meu	
		nome
		é
		Dihego
	`
	fmt.Println(s2)
	fmt.Println("O tamanho da string é", len(s2))

}
