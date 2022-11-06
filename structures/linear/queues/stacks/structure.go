package stacks

import (
	"github.com/denstepanov/algostructs/structures/linear/stacks/slice"
)

type StacksQueue[T comparable] struct {
	in, out slice.SliceStack[T]
}

func (q *StacksQueue[T]) IsEmpty() bool {
	return q.in.Len() == 0 && q.out.Len() == 0
}

func (q *StacksQueue[T]) Len() int {
	return q.in.Len() + q.out.Len()
}

func (q *StacksQueue[T]) Enqueue(item T) {
	if !q.out.IsEmpty() {
		q.fromOutToIn()
	}
	q.in.Push(item)
}

func (q *StacksQueue[T]) Dequeue() (item T) {
	if q.IsEmpty() {
		return item
	}

	if q.in.Len() != 0 {
		q.fromInToOut()
	}

	return q.out.Pop()
}

func (q *StacksQueue[T]) Peek() (item T) {
	if q.IsEmpty() {
		return item
	}

	if q.in.Len() != 0 {
		q.fromInToOut()
	}

	item = q.out.Pop()
	q.out.Push(item)
	return item
}

func (q *StacksQueue[T]) fromInToOut() {
	for !q.in.IsEmpty() {
		item := q.in.Pop()
		q.out.Push(item)
	}
}

func (q *StacksQueue[T]) fromOutToIn() {
	for !q.out.IsEmpty() {
		item := q.out.Pop()
		q.in.Push(item)
	}
}
