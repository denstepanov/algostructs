package queues

import (
	"errors"

	"github.com/denstepanov/algostructs/structures"
	"github.com/denstepanov/algostructs/structures/linear/stacks"
)

type StacksQueue[T comparable] struct {
	in, out stacks.StackViaSlice[T]
}

func (q *StacksQueue[T]) IsEmpty() bool {
	return q.in.Len() == 0 && q.out.Len() == 0
}

func (q *StacksQueue[T]) Len() int {
	return q.in.Len() + q.out.Len()
}

// Вставка элемента в начало.
func (q *StacksQueue[T]) Enqueue(item T) {
	if !q.out.IsEmpty() {
		q.fromOutToIn()
	}
	q.in.Push(item)
}

// Удаление элемента с конца.
func (q *StacksQueue[T]) Dequeue() (item T, err error) {
	if q.IsEmpty() {
		err = errors.New(structures.EmptyQueue)
		return item, err
	}

	if q.in.Len() != 0 {
		q.fromInToOut()
	}

	return q.out.Pop()
}

// Возврат последнего элемента.
func (q *StacksQueue[T]) Peek() (item T, err error) {
	if q.IsEmpty() {
		err = errors.New(structures.EmptyQueue)
		return item, err
	}

	if q.in.Len() != 0 {
		q.fromInToOut()
	}

	item, _ = q.out.Pop()
	q.out.Push(item)
	return item, nil
}

func (q *StacksQueue[T]) fromInToOut() {
	for !q.in.IsEmpty() {
		item, _ := q.in.Pop()
		q.out.Push(item)
	}
}

func (q *StacksQueue[T]) fromOutToIn() {
	for !q.out.IsEmpty() {
		item, _ := q.out.Pop()
		q.in.Push(item)
	}
}
