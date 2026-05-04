package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Teste 1: Básico (2 + 1) * 3 = 9
	teste1 := []string{"2", "1", "+", "3", "*"}
	res1 := evalRPN(teste1)
	fmt.Printf("Teste 1: %v | Resultado: %d (Esperado: 9)\n", teste1, res1)

	// Teste 2: Ordem da divisão e números negativos
	// (4 + (13 / 5)) = 4 + 2 = 6
	teste2 := []string{"4", "13", "5", "/", "+"}
	res2 := evalRPN(teste2)
	fmt.Printf("Teste 2: %v | Resultado: %d (Esperado: 6)\n", teste2, res2)

	// Teste 3: Expressão complexa
	// ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
	// O LeetCode usa muito esse tipo de caso para testar performance
	teste3 := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	res3 := evalRPN(teste3)
	fmt.Printf("Teste 3: %v | Resultado: %d (Esperado: 22)\n", teste3, res3)
}

type StackNode struct {
	Val  int
	Next *StackNode
}

type Stack struct {
	Topo *StackNode
	Size int
}

func (s *Stack) Push(v int) {
	newNode := &StackNode{Val: v}
	newNode.Next = s.Topo
	s.Topo = newNode
	s.Size++
}

func (s *Stack) Pop() (int, bool) {
	if s.Size == 0 || s.Topo == nil {
		return 0, false
	}
	item := s.Topo.Val
	s.Topo = s.Topo.Next
	s.Size--
	return item, true
}

func (s *Stack) Peek() (int, bool) {
	if s.Topo == nil {
		return 0, false
	}
	return s.Topo.Val, true
}

func evalRPN(tokens []string) int {
	stack := &Stack{}

	for _, item := range tokens {
		if item == "/" || item == "*" || item == "+" || item == "-" {
			num2, _ := stack.Pop()
			num1, _ := stack.Pop()

			switch item {
			case "+":
				stack.Push(num1 + num2)
			case "-":
				stack.Push(num1 - num2)
			case "*":
				stack.Push(num1 * num2)
			case "/":
				stack.Push(num1 / num2)
			}
		} else {
			valorInt, _ := strconv.Atoi(item)
			stack.Push(valorInt)
		}
	}
	resultado, _ := stack.Pop()
	return resultado
}
