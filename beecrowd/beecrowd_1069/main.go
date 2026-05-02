package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Scanln()

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		if scanner.Scan() {
			linha := scanner.Text()
			fmt.Println(checkDiam(linha))
		}
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

func checkDiam(s string) int {
	stack := &Stack{}
	qtdDiamantes := 0
	if len(s) == 0 {
		return 0
	}
	for _, char := range s {
		if char == '<' {
			stack.Push(char)
		} else if char == '>' {
			itemPeek, ok := stack.Peek()
			if ok && itemPeek == '<' {
				stack.Pop()
				qtdDiamantes++
			}
		}
	}
	return qtdDiamantes
}
