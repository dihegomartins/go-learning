package main

import "fmt"

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
	itemValue := s.Topo.Val
	s.Topo = s.Topo.Next
	s.Size--
	return itemValue, true
}

func (s *Stack) Peek() (rune, bool) {
	if s.Topo == nil {
		return 0, false
	}
	return s.Topo.Val, true
}

func isValid(s string) bool {
	stack := &Stack{}
	maps := make(map[rune]rune)
	maps[']'] = '['
	maps[')'] = '('
	maps['}'] = '{'

	if len(s) == 0 {
		return true
	}
	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack.Push(char)
		} else if itemPeek, ok := stack.Peek(); ok && itemPeek == maps[char] {
			stack.Pop()
		} else {
			return false
		}
	}
	return stack.Size == 0
}

func main() {
	// Casos de teste
	testes := []string{
		"()",     // Verdadeiro
		"()[]{}", // Verdadeiro
		"(]",     // Falso
		"([)]",   // Falso
		"{[]}",   // Verdadeiro
		"[",      // Falso (ajuda a testar o return stack.Size == 0)
		"]",      // Falso (ajuda a testar o !existe no Peek)
	}

	fmt.Println("--- Testando Pilha de Parênteses ---")

	for _, t := range testes {
		resultado := isValid(t)
		status := "❌ FALHOU"
		if (resultado && (t == "()" || t == "()[]{}" || t == "{[]}")) ||
			(!resultado && (t == "(]" || t == "([)]" || t == "[" || t == "]")) {
			status = "✅ PASSOU"
		}

		fmt.Printf("String: %-8s | Resultado: %-5v | %s\n", t, resultado, status)
	}
}
