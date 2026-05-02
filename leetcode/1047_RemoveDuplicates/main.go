package main

import (
	"fmt"
)

func main() {
	testes := []struct {
		entrada  string
		esperado string
	}{
		{"abbaca", "ca"},     // 'bb' cancela, sobra 'aaca', 'aa' cancela, sobra 'ca'
		{"azxxzy", "ay"},     // 'xx' cancela, depois 'zz' cancela, sobra 'ay'
		{"aaaaaaaa", ""},     // Todos se cancelam em pares
		{"abcdef", "abcdef"}, // Ninguém se cancela
		{"", ""},             // Caso vazio
	}

	fmt.Println("--- Testando Remoção de Duplicatas Adjacentes ---")

	for _, t := range testes {
		resultado := removeDuplicates(t.entrada) // Nome da função padrão do LeetCode

		status := "❌ FALHOU"
		if resultado == t.esperado {
			status = "✅ PASSOU"
		}

		fmt.Printf("Entrada: %-10s | Saída: %-10s | Esperado: %-10s | %s\n",
			t.entrada, resultado, t.esperado, status)
	}
}

type StackNode struct {
	Val  rune
	Next *StackNode
}

type Stack struct {
	Topo *StackNode
	Size int
}

func (s *Stack) Push(v rune) {
	newNode := &StackNode{Val: v}
	newNode.Next = s.Topo
	s.Topo = newNode
	s.Size++
}

func (s *Stack) Pop() (rune, bool) {
	if s.Size == 0 || s.Topo == nil {
		return 0, false
	}
	item := s.Topo.Val
	s.Topo = s.Topo.Next
	s.Size--
	return item, true
}

func (s *Stack) Peek() (rune, bool) {
	if s.Topo == nil {
		return 0, false
	}

	return s.Topo.Val, true
}

func removeDuplicates(s string) string {
	stack := &Stack{}

	if len(s) == 0 {
		return s
	}
	for _, char := range s {
		itemPeek, ok := stack.Peek()

		if ok && itemPeek == char {
			stack.Pop()
		} else {
			stack.Push(char)
		}
	}

	res := make([]rune, stack.Size)
	for i := stack.Size - 1; i >= 0; i-- {
		val, _ := stack.Pop()
		res[i] = val
	}
	return string(res)
}
