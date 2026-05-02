package main

import "fmt"

func main() {
	// Criando os nós manualmente
	n4 := &ListNode{Val: 40, Next: nil}
	n3 := &ListNode{Val: 30, Next: n4}
	n2 := &ListNode{Val: 20, Next: n3}
	n1 := &ListNode{Val: 10, Next: n2}

	fmt.Println("Original:")
	printList(n1)

	// Testando a reversão
	reversed := reverseList(n1)
	fmt.Println("Invertida:")
	printList(reversed)
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

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := head

	// [10] -> [20] -> [30] -> nil.

	// nill <- [10] <- [20]

	for curr != nil {
		next := curr.Next // next = 20			// next = 30
		curr.Next = prev  // curr.Next = nil	//curr.Next = O nó 20->10. (O ponteiro virou para trás!).
		prev = curr       // prev = 10			// prev = 20
		curr = next       // curr = 20			// curr = 30
	}
	return prev
}

func printList(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Printf("[%d] -> ", curr.Val)
		curr = curr.Next
	}
	fmt.Println("nil")
}
