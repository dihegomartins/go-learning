package main

import "fmt"

func main() {
	// 1. Criando os nós isolados
	node1 := &ListNode{Val: 5}
	node2 := &ListNode{Val: 6}
	node3 := &ListNode{Val: 9}
	node4 := &ListNode{Val: 5}
	node5 := &ListNode{Val: 1}

	// 2. Conectando os nós linearmente: 5 -> 6 -> 9 -> 5 -> 1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	// 3. CRIANDO O CICLO (pos = 0)
	// Aqui fazemos o último nó (1) apontar de volta para o primeiro nó (5)
	node5.Next = node1

	// Chamada da sua função
	resultado := hasCycle(node1)

	if resultado {
		fmt.Println("Resultado: true (Ciclo detectado!)")
	} else {
		fmt.Println("Resultado: false (Sem ciclo)")
	}
}

/**
 * Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
