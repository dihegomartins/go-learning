package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "1 + 1"
	resultado := calculate(s)
	fmt.Println(resultado)
}

type StackNode[T any] struct {
	Val  T
	Next *StackNode[T]
}

type Stack[T any] struct {
	Topo *StackNode[T]
	Size int
}

func (s *Stack[T]) Push(v T) {
	newNode := &StackNode[T]{Val: v}
	newNode.Next = s.Topo
	s.Topo = newNode
	s.Size++
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.Topo == nil {
		var zero T
		return zero, false
	}
	item := s.Topo.Val
	s.Topo = s.Topo.Next
	s.Size--
	return item, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.Topo == nil {
		var zero T
		return zero, false
	}
	return s.Topo.Val, true
}

func calculate(s string) int {
	stack := &Stack[int]{}

	s = strings.ReplaceAll(s, " ", "")
	resultado := 0
	num := 0
	sinal := 1

	for _, item := range s {
		if item >= '0' && item <= '9' {
			num = num*10 + int(item-'0')
		} else if item == '+' || item == '-' {
			resultado += sinal * num
			num = 0
			if item == '-' {
				sinal = -1
			} else {
				sinal = 1
			}
		} else if item == '(' {
			stack.Push(resultado)
			stack.Push(sinal)
			resultado = 0
			sinal = 1
			num = 0
		} else if item == ')' {
			resultado += sinal * num
			num = 0
			sinalAnterior, _ := stack.Pop()
			resultado *= sinalAnterior

			resultadoAnterior, _ := stack.Pop()
			resultado += resultadoAnterior
		}
	}
	return resultado + (sinal * num)
}
