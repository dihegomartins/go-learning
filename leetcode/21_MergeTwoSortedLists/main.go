package main

import "fmt"

func main() {
	// Criando a primeira lista ordenada: 1 -> 3 -> 5
	l1 := &ListNode{Val: 1}
	l1.Next = &ListNode{Val: 3}
	l1.Next.Next = &ListNode{Val: 5}

	// Criando a segunda lista ordenada: 2 -> 4 -> 6
	l2 := &ListNode{Val: 2}
	l2.Next = &ListNode{Val: 4}
	l2.Next.Next = &ListNode{Val: 6}

	// Chamando a função que você vai criar
	// Ela deve retornar a cabeça da nova lista mesclada
	resultado := mergeTwoLists(l1, l2)

	// Print para conferir o resultado
	fmt.Print("Lista Mesclada: ")
	curr := resultado
	for curr != nil {
		fmt.Printf("%d -> ", curr.Val)
		curr = curr.Next
	}
	fmt.Println("nil")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	DummyNode := &ListNode{Val: 0}

	curr := DummyNode

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}
	if list1 == nil {
		curr.Next = list2
	} else if list2 == nil {
		curr.Next = list1
	}

	return DummyNode.Next
}
