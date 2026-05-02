package main

import "fmt"

func main() {
	a := []int32{2, 4}
	b := []int32{24, 36}

	resultado := getTotalX(a, b)
	fmt.Println("Resultado esperado: 2 | Resultado obtido:", resultado)
}

func getTotalX(a []int32, b []int32) int32 {
	minimoComum := algoritmoMMC(a)
	maximoDivisor := MDCdMultiple(b)

	var contador int32 = 0

	for i := minimoComum; i <= maximoDivisor; i += minimoComum {
		if maximoDivisor%i == 0 {
			contador++
		}
	}
	return contador
}

func algoritmosEuclidesMDC(a, b int32) int32 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func MDCdMultiple(b []int32) int32 {
	result := b[0]
	for i := 1; i < len(b); i++ {
		result = algoritmosEuclidesMDC(result, b[i])
	}
	return result
}

func MMC(a, b int32) int32 {
	if a == 0 || b == 0 {
		return 0
	}

	return (a * b) / algoritmosEuclidesMDC(a, b)
}

func algoritmoMMC(a []int32) int32 {
	if len(a) == 0 {
		return 0
	}
	result := a[0]
	for i := 1; i < len(a); i++ {
		result = MMC(result, a[i])
	}
	return result
}

/*
   Complexidade de Tempo: O(n log(max) + m log(max) + (MDC/MMC))
   - 'n' e 'm' são os tamanhos dos arrays A e B.
   - O log(max) vem do algoritmo de Euclides (MDC).
   - O termo final (MDC/MMC) é o número de iterações do loop final.
   No geral, é muito mais eficiente que testar todos os números no intervalo.

   Complexidade de Espaço: O(1)
   Usamos apenas variáveis simples para armazenar os resultados dos cálculos,
   sem criar estruturas de dados que crescem com a entrada.
*/
