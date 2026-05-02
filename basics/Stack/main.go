package main

func main() {

}

type StackNode struct {
	Val  int
	Next *StackNode
}

type Stack struct {
	Topo *StackNode
	Size int
}

func (q *Stack) Push(v int) {
	newNode := &StackNode{Val: v}
	newNode.Next = q.Topo
	q.Topo = newNode

	q.Size++
}

func (q *Stack) Pop() (int, bool) {
	if q.Topo == nil {
		return 0, false
	}
	item := q.Topo.Val
	q.Topo = q.Topo.Next
	q.Size--
	return item, true
}
