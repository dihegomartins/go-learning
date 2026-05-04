package main

import (
	"strconv"
	"strings"
)

func main() {

}

type StackNode struct {
	Val  int32
	Next *StackNode
	Max  int32
}

type MinStack struct {
	Topo *StackNode
	Size int32
}

func Constructor() MinStack {
	return MinStack{
		Topo: nil,
		Size: 0,
	}
}

func (this *MinStack) Push(val int32) {
	newNode := &StackNode{Val: val, Max: val}
	newNode.Next = this.Topo
	this.Topo = newNode

	if this.Topo != nil && this.Topo.Next != nil {
		if this.Topo.Max < this.Topo.Next.Max {
			newNode.Max = this.Topo.Next.Max
		}
	}

	this.Size++
}

func (this *MinStack) Pop() {
	if this.Topo != nil {
		this.Topo = this.Topo.Next
		this.Size--
	}
}

func (this *MinStack) Top() int32 {
	if this.Topo == nil {
		return 0
	}
	return this.Topo.Val
}

func getMax(operations []string) []int32 {
	stack := &MinStack{}
	var resultados []int32

	for _, op := range operations {
		tipoOperacao := op[0]

		switch tipoOperacao {
		case '1':
			partes := strings.Split(op, " ")
			valorInt, _ := strconv.Atoi(partes[1])
			stack.Push(int32(valorInt))
		case '2':
			stack.Pop()
		case '3':
			resultados = append(resultados, stack.Topo.Max)
		}
	}
	return resultados
}
