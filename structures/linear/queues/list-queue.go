package queues

import (
	"github.com/denstepanov/algostructs/structures/linear/lists"
)

type ListQueue[T comparable] struct {
	list lists.DLList[T]
}

func (q *ListQueue[T]) IsEmpty() bool {
	return q.list.IsEmpty()
}

func (q *ListQueue[T]) Len() int {
	return q.list.Len()
}

func (q *ListQueue[T]) Enqueue(item *T) {
	node := &lists.DLNode[T]{
		Value: item,
	}
	q.list.InsertHead(node)
}

func (q *ListQueue[T]) Dequeue() *T {
	if q.list.IsEmpty() {
		return nil
	}

	return q.list.DeleteTail().Value
}

func (q *ListQueue[T]) Peek() *T {
	if q.list.IsEmpty() {
		return nil
	}

	tail := q.list.DeleteTail()
	q.list.InsertTail(tail)
	return tail.Value
}
