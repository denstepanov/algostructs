package queues

import (
	"errors"

	"github.com/denstepanov/algostructs/structures"
)

type SliceQueue[T comparable] []T

func (q *SliceQueue[T]) IsEmpty() bool {
	return len(*q) == 0
}

func (q *SliceQueue[T]) Len() int {
	return len(*q)
}

func (q *SliceQueue[T]) Enqueue(item T) {
	newQueue := []T{item}
	if !q.IsEmpty() {
		newQueue = append(newQueue, *q...)
	}
	*q = newQueue
}

func (q *SliceQueue[T]) Dequeue() (item T, err error) {
	if q.IsEmpty() {
		err = errors.New(structures.EmptyQueue)
		return item, err
	}
	index := len(*q) - 1
	item = (*q)[index]
	*q = append([]T{}, (*q)[:index]...)
	return item, nil
}

func (q *SliceQueue[T]) Peek() (item T, err error) {
	if q.IsEmpty() {
		err = errors.New(structures.EmptyQueue)
		return item, err
	}
	return (*q)[len(*q)-1], nil
}
