package main

func main() {

}

type StackNode struct {
	Val  int
	Next *StackNode
	Min  int
}

type MinStack struct {
	Topo *StackNode
	Size int
}

func Constructor() MinStack {
	return MinStack{
		Topo: nil,
		Size: 0,
	}
}

func (this *MinStack) Push(val int) {
	newNode := &StackNode{Val: val, Min: val}
	newNode.Next = this.Topo
	this.Topo = newNode

	if this.Topo != nil && this.Topo.Next != nil {
		if this.Topo.Min > this.Topo.Next.Min {
			newNode.Min = this.Topo.Next.Min
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

func (this *MinStack) Top() int {
	if this.Topo == nil {
		return 0
	}
	return this.Topo.Val
}

func (this *MinStack) GetMin() int {
	return this.Topo.Min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
