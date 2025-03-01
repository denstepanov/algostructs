package linked_lists

import (
	"github.com/denstepanov/algostructs/src/structures/linear/linked-lists/doubly"
)

type Queue[T comparable] struct {
	list doubly.LinkedList[T]
}

func New[T comparable]() *Queue[T] {
	return new(Queue[T])
}

func (q *Queue[T]) IsEmpty() bool {
	return q.list.IsEmpty()
}

func (q *Queue[T]) Len() int {
	return q.list.Len()
}

func (q *Queue[T]) Enqueue(item T) {
	node := &doubly.Node[T]{
		Value: item,
	}
	q.list.InsertHead(node)
}

func (q *Queue[T]) Dequeue() T {
	var result T
	if !q.list.IsEmpty() {
		result = q.list.DeleteTail().Value
	}
	return result
}

func (q *Queue[T]) Peek() T {
	var result T
	if !q.list.IsEmpty() {
		tail := q.list.DeleteTail()
		q.list.InsertTail(tail)
		result = tail.Value
	}
	return result
}
