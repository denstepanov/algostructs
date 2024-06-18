package stacks

import (
	"github.com/denstepanov/algostructs/src/structure/linear/stack/slice"
)

type Queue[T comparable] struct {
	in, out slice.Stack[T]
}

func New[T comparable]() *Queue[T] {
	return new(Queue[T])
}

func (q *Queue[T]) Len() int {
	return q.in.Len() + q.out.Len()
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Len() == 0
}

func (q *Queue[T]) Enqueue(item T) {
	if !q.out.IsEmpty() {
		q.fromOutToIn()
	}
	q.in.Push(item)
}

func (q *Queue[T]) Dequeue() T {
	var result T
	if !q.IsEmpty() {
		if q.in.Len() != 0 {
			q.fromInToOut()
		}
		result = q.out.Pop()
	}
	return result
}

func (q *Queue[T]) Peek() T {
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

func (q *Queue[T]) fromInToOut() {
	for !q.in.IsEmpty() {
		item := q.in.Pop()
		q.out.Push(item)
	}
}

func (q *Queue[T]) fromOutToIn() {
	for !q.out.IsEmpty() {
		item := q.out.Pop()
		q.in.Push(item)
	}
}
