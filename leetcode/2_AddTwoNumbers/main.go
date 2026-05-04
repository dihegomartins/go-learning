package main

import (
	"fmt"
)

func main() {
	// Criando as listas manualmente para bater com o exemplo da image_d1a582.png
	// L1: 2 -> 4 -> 3
	node3 := &ListNode{Val: 3}
	node2 := &ListNode{Val: 4, Next: node3}
	l1 := &ListNode{Val: 2, Next: node2}

	// L2: 5 -> 6 -> 4
	node6 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 6, Next: node6}
	l2 := &ListNode{Val: 5, Next: node5}

	res := addTwoNumbers(l1, l2)

	fmt.Print("Resultado do seu código: ")
	for res != nil {
		fmt.Printf("%d ", res.Val)
		res = res.Next
	}
	// Com esse código, o resultado será 7 0 8 porque 243 + 564 = 807,
	// e o Add(8), Add(0), Add(7) inverte para 7 0 8.
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type List struct {
	Head *ListNode
	Size int
}

func (s *List) Add(v int) {
	newNode := &ListNode{Val: v}
	newNode.Next = s.Head
	s.Head = newNode
	s.Size++
}

func (s *List) RemoveHead() (int, bool) {
	if s.Head == nil {
		return 0, false
	}
	item := s.Head.Val
	s.Head = s.Head.Next
	s.Size--
	return item, true
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	list1 := &List{}
	curr := l1
	curr2 := l2
	carry := 0

	for curr != nil || curr2 != nil || carry > 0 {
		v1, v2 := 0, 0

		if curr != nil {
			v1 = curr.Val
			curr = curr.Next
		} else {
			v1 = 0
		}
		if curr2 != nil {
			v2 = curr2.Val
			curr2 = curr2.Next
		} else {
			v2 = 0
		}

		sum := v1 + v2 + carry
		carry = sum / 10

		list1.Add(sum % 10)
	}

	return list1.Head
}
