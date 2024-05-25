package stacks

import (
	"github.com/denstepanov/algostructs/src/structures/linear/stacks/slice"
)

type StacksQueue[T comparable] struct {
	in, out slice.SliceStack[T]
}

func New[T comparable]() *StacksQueue[T] {
	return new(StacksQueue[T])
}

func (q *StacksQueue[T]) Len() int {
	return q.in.Len() + q.out.Len()
}

func (q *StacksQueue[T]) IsEmpty() bool {
	return q.Len() == 0
}

func (q *StacksQueue[T]) Enqueue(item T) {
	if !q.out.IsEmpty() {
		q.fromOutToIn()
	}
	q.in.Push(item)
}

func (q *StacksQueue[T]) Dequeue() T {
	var result T
	if !q.IsEmpty() {
		if q.in.Len() != 0 {
			q.fromInToOut()
		}
		result = q.out.Pop()
	}
	return result
}

func (q *StacksQueue[T]) Peek() T {
	var result T
	if !q.IsEmpty() {
		if q.in.Len() != 0 {
			q.fromInToOut()
		}
		result = q.out.Pop()
		q.out.Push(result)
	}
	return result
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
