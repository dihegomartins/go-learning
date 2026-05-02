package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type MyLinkedList struct {
	head *ListNode
	size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		head: nil,
		size: 0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	} else {
		curr := this.head
		for i := 0; i < index; i++ {
			curr = curr.Next
		}
		return curr.Val
	}
}

func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &ListNode{Val: val}
	newNode.Next = this.head

	this.head = newNode
	this.size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	if this.size == 0 {
		this.AddAtHead(val)
	} else {
		newNode := &ListNode{Val: val}

		curr := this.head
		for curr.Next != nil {
			curr = curr.Next
		}
		curr.Next = newNode
		this.size++
	}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size {
		return
	}
	if index == this.size {
		this.AddAtTail(val)
	} else if index == 0 {
		this.AddAtHead(val)
	} else {
		newNode := &ListNode{Val: val}

		curr := this.head
		for i := 0; i < index-1; i++ {
			curr = curr.Next
		}
		newNode.Next = curr.Next
		curr.Next = newNode
		this.size++
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}
	if index == 0 {
		if this.head != nil {
			this.head = this.head.Next
		}
	} else {
		curr := this.head

		for i := 0; i < index-1; i++ {
			curr = curr.Next
		}

		if curr != nil {
			curr.Next = curr.Next.Next
		}
	}
	this.size--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
