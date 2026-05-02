package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		linha := scanner.Text()
		if isValid(linha) {
			fmt.Println("correct")
		} else {
			fmt.Println("incorrect")
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
	maps[')'] = '('

	if len(s) == 0 {
		return true
	}
	for _, char := range s {
		if char == '(' {
			stack.Push(char)
		} else if char == ')' {
			itemPeek, ok := stack.Peek()

			if ok && itemPeek == '(' {
				stack.Pop()
			} else {
				return false
			}
		}
	}
	return stack.Size == 0
}
