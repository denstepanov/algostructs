package list

import (
	"github.com/denstepanov/algostructs/structures/linear/lists/doubly"
)

type ListQueue[T comparable] struct {
	list doubly.List[T]
}

func New[T comparable]() *ListQueue[T] {
	return new(ListQueue[T])
}

func (q *ListQueue[T]) IsEmpty() bool {
	return q.list.IsEmpty()
}

func (q *ListQueue[T]) Len() int {
	return q.list.Len()
}

func (q *ListQueue[T]) Enqueue(item T) {
	node := &doubly.Node[T]{
		Value: item,
	}
	q.list.InsertHead(node)
}

func (q *ListQueue[T]) Dequeue() T {
	var result T
	if !q.list.IsEmpty() {
		result = q.list.DeleteTail().Value
	}
	return result
}

func (q *ListQueue[T]) Peek() T {
	var result T
	if !q.list.IsEmpty() {
		tail := q.list.DeleteTail()
		q.list.InsertTail(tail)
		result = tail.Value
	}
	return result
}
