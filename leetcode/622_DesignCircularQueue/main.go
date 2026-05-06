package main

func main() {

}

type QueueNode struct {
	Val  int
	Next *QueueNode
}

type MyCircularQueue struct {
	Frontt     *QueueNode
	Tail       *QueueNode
	Capacidade int
	Size       int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		Frontt:     nil,
		Tail:       nil,
		Capacidade: k,
		Size:       0,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.Size == this.Capacidade {
		return false
	}
	newNode := &QueueNode{Val: value}
	if this.Size == 0 {
		this.Frontt = newNode
		this.Tail = newNode
	} else {
		this.Tail.Next = newNode
		this.Tail = newNode
	}
	this.Size++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.Size == 0 {
		return false
	}

	this.Frontt = this.Frontt.Next
	this.Size--

	if this.Size == 0 {
		this.Tail = this.Frontt
	}
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.Size == 0 {
		return -1
	}
	return this.Frontt.Val
}

func (this *MyCircularQueue) Rear() int {
	if this.Size == 0 {
		return -1
	}
	return this.Tail.Val
}

func (this *MyCircularQueue) IsEmpty() bool {
	if this.Size == 0 {
		return true
	}
	return false
}

func (this *MyCircularQueue) IsFull() bool {
	if this.Size == this.Capacidade {
		return true
	}
	return false
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
