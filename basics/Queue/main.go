package main

func main() {

}

type QueueNode[T any] struct {
	Val  T
	Next *QueueNode[T]
}

type Queue[T any] struct {
	Front *QueueNode[T]
	Tail  *QueueNode[T]
	Size  int
}

func (q *Queue[T]) Enqueue(v T) {
	newNode := &QueueNode[T]{Val: v}
	if q.Size == 0 {
		q.Front = newNode
		q.Tail = newNode
	} else {
		q.Tail.Next = newNode
		q.Tail = newNode
	}
	q.Size++
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if q.Size == 0 {
		var zero T
		return zero, false
	}
	item := q.Front.Val
	q.Front = q.Front.Next
	q.Size--
	if q.Size == 0 {
		q.Tail = q.Front
	}
	return item, true
}

func (q *Queue[T]) First() (T, bool) {
	if q.Front == nil {
		var zero T
		return zero, false
	}
	return q.Front.Val, true
}
