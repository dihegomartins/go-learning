package main

import "fmt"

func main() {
	testes := []struct {
		s        string
		t        string
		esperado bool
	}{
		{"ab#c", "ad#c", true}, // Ambos viram "ac"
		{"ab##", "c#d#", true}, // Ambos viram ""
		{"a#c", "b", false},    // Um vira "c", o outro é "b"
		{"a##c", "#a#c", true}, // Ambos viram "c" (backspace no vazio é ignorado)
		{"at#c#", "a", true},   // "at#c#" vira "a"
	}

	fmt.Println("--- Testando Backspace String Compare ---")

	for _, teste := range testes {
		resultado := backspaceCompare(teste.s, teste.t)

		status := "❌ FALHOU"
		if resultado == teste.esperado {
			status = "✅ PASSOU"
		}

		fmt.Printf("S: %-7s | T: %-7s | Resultado: %-5v | Esperado: %-5v | %s\n",
			teste.s, teste.t, resultado, teste.esperado, status)
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

func RemoveLetras(palavra string) string {
	stack := &Stack{}
	for _, char := range palavra {
		if char != '#' {
			stack.Push(char)
		} else {
			stack.Pop()
		}
	}
	minhaPalavra := ""
	for stack.Size > 0 {
		letra, ok := stack.Pop()
		if ok {
			minhaPalavra = string(letra) + minhaPalavra
		}
	}
	return minhaPalavra
}

func backspaceCompare(s string, t string) bool {
	s = RemoveLetras(s)
	t = RemoveLetras(t)
	if s == t {
		return true
	}
	return false
}
