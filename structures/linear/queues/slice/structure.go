package slice

type SliceQueue[T comparable] []T

func New[T comparable]() *SliceQueue[T] {
	return new(SliceQueue[T])
}

func (q *SliceQueue[T]) Len() int {
	return len(*q)
}

func (q *SliceQueue[T]) IsEmpty() bool {
	return q.Len() == 0
}

func (q *SliceQueue[T]) Enqueue(item T) {
	newQueue := []T{item}
	if !q.IsEmpty() {
		newQueue = append(newQueue, *q...)
	}
	*q = newQueue
}

func (q *SliceQueue[T]) Dequeue() T {
	var result T
	if !q.IsEmpty() {
		index := len(*q) - 1
		result = (*q)[index]
		*q = append([]T{}, (*q)[:index]...)
	}
	return result
}

func (q *SliceQueue[T]) Peek() T {
	var result T
	if !q.IsEmpty() {
		result = (*q)[len(*q)-1]
	}
	return result
}
