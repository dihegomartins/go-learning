package main

import (
	"errors"
	"fmt"
)

func main() {
	minhaLista := Constructor()
	minhaLista.Add(10.5)
	minhaLista.Add(6.2)
	minhaLista.Add(5.2)
	minhaLista.Add(3.2)
	minhaLista.Add(1.6)
	minhaLista.Add(2.9)
	minhaLista.Add(3.9)

	toString(minhaLista.head)
	minhaLista.BubbleSort()
	toString(minhaLista.head)

}

type Node struct {
	Val  float64
	Next *Node
}

type List struct {
	head *Node
	size int
}

func Constructor() List {
	return List{
		head: nil,
		size: 0,
	}
}

func (this *List) Add(val float64) {
	newNode := &Node{Val: val}

	newNode.Next = this.head
	this.head = newNode
	this.size++
}

func (this *List) BubbleSort() {

	houveTroca := true
	for houveTroca {
		curr := this.head

		houveTroca = false
		for curr.Next != nil {
			if curr.Val > curr.Next.Val {
				curr.Val, curr.Next.Val = curr.Next.Val, curr.Val
				houveTroca = true
			}
			curr = curr.Next
		}
	}

}

func (this *List) RemoveHead() (float64, error) {
	if this.head == nil {
		return 0, errors.New("A lista esta vazia!")
	}
	item := this.head.Val
	this.head = this.head.Next
	this.size--
	return item, nil
}

func (this *List) RemoveManys(n int) error {
	for i := 0; i < n; i++ {
		_, err := this.RemoveHead()
		if err != nil {
			return err
		}
	}
	return nil
}

func (this *List) RemoveLastNumber() (float64, error) {
	if this.head == nil {
		return 0, errors.New("Lista vazia")
	} else if this.head.Next == nil {
		item := this.head.Val
		this.RemoveHead()
		return item, nil
	}
	curr := this.head
	for curr.Next.Next != nil {
		curr = curr.Next
	}

	if curr.Next == nil {
		return 0, errors.New("erro inesperado: o último nó desapareceu")
	}

	itemRem := curr.Next.Val
	curr.Next = nil
	this.size--
	return itemRem, nil
}

func toString(head *Node) {
	curr := head
	size := 0
	fmt.Printf("[")
	for curr != nil {
		fmt.Printf("%.2f -> ", curr.Val)
		curr = curr.Next
		size++
	}
	fmt.Println("nil]")
	fmt.Println("Tamanho da Lista:", size)
}
