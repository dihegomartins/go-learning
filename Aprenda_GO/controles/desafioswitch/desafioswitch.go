package main

import "fmt"

func main() {
	fmt.Println(notaParaConceito(10))
	fmt.Println(notaParaConceito(9))
	fmt.Println(notaParaConceito(8))
	fmt.Println(notaParaConceito(1))
	fmt.Println(notaParaConceito(11))
}

func notaParaConceito(n float64) string {
	switch {
	case n >= 9 && n <= 10:
		return "A"
	case n >= 8 && n < 9:
		return "B"
	case n >= 5 && n < 8:
		return "C"
	case n >= 3 && n < 5:
		return "D"
	case n >= 0 && n < 3:
		return "E"
	default:
		return "Nota inválida"
	}
}
