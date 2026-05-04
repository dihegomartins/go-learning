package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var linha string
		fmt.Scan(&linha)
		inputSlice := strings.Split(linha, "")

		fmt.Println(infixToPostfix(inputSlice))
	}
}

type StackNode struct {
	Val  string
	Next *StackNode
}

type Stack struct {
	Topo *StackNode
	Size int
}

func (s *Stack) Push(v string) {
	newNode := &StackNode{Val: v}
	newNode.Next = s.Topo
	s.Topo = newNode
	s.Size++
}

func (s *Stack) Pop() (string, bool) {
	if s.Topo == nil {
		return "", false
	}
	item := s.Topo.Val
	s.Topo = s.Topo.Next
	s.Size--
	return item, true
}

func (s *Stack) Peek() (string, bool) {
	if s.Topo == nil {
		return "", false
	}
	return s.Topo.Val, true
}

func precedencia(op string) int {
	switch op {
	case "^":
		return 3
	case "*", "/":
		return 2
	case "+", "-":
		return 1
	}
	return 0
}

func infixToPostfix(equacao []string) string {
	stackOp := &Stack{}
	resultado := ""
	for _, item := range equacao {
		if item == "(" {
			stackOp.Push(item)
		} else if item == ")" {
			for {
				op, _ := stackOp.Pop()
				if op == "(" {
					break
				}
				resultado += op
			}
		} else if item == "+" || item == "-" || item == "*" || item == "/" || item == "^" {
			for stackOp.Size > 0 && precedencia(stackOp.Topo.Val) >= precedencia(item) {
				op, _ := stackOp.Pop()
				resultado += op
			}
			stackOp.Push(item)
		} else {
			resultado += item
		}
	}
	for stackOp.Size > 0 {
		op, _ := stackOp.Pop()
		resultado += op
	}
	return resultado
}
