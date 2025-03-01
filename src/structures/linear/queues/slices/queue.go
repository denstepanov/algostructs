package slices

type Queue[T comparable] []T

func New[T comparable]() *Queue[T] {
	return new(Queue[T])
}

func (q *Queue[T]) Len() int {
	return len(*q)
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Len() == 0
}

func (q *Queue[T]) Enqueue(item T) {
	newQueue := []T{item}
	if !q.IsEmpty() {
		newQueue = append(newQueue, *q...)
	}
	*q = newQueue
}

func (q *Queue[T]) Dequeue() T {
	var result T
	if !q.IsEmpty() {
		index := len(*q) - 1
		result = (*q)[index]
		*q = append([]T{}, (*q)[:index]...)
	}
	return result
}

func (q *Queue[T]) Peek() T {
	var result T
	if !q.IsEmpty() {
		result = (*q)[len(*q)-1]
	}
	return result
}
