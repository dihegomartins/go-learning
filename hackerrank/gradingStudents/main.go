package main

import "fmt"

func main() {
	ar := []int32{84, 57, 37, 38, 75}
	fmt.Println(gradingStudents(ar))
}

func gradingStudents(grades []int32) []int32 {
	for i := 0; i < len(grades); i++ {
		if resto := grades[i] % 5; grades[i] >= 38 && 5-resto < 3 {
			grades[i] = grades[i] + (5 - resto)
		}
	}
	return grades
}

/*
   Complexidade de Tempo: O(n) - Percorre o array 'grades' uma única vez.
   Complexidade de Espaço: O(1) - Modifica os valores in-place, sem alocação extra.
*/
