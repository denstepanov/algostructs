package slice

type SliceStack[T comparable] []T

func (s *SliceStack[T]) IsEmpty() bool {
	return len(*s) == 0
}

func (s *SliceStack[T]) Len() int {
	return len(*s)
}

func (s *SliceStack[T]) Push(item T) {
	*s = append(*s, item)
}

func (s *SliceStack[T]) Pop() (item T) {
	if s.IsEmpty() {
		return item
	}
	index := len(*s) - 1
	item = (*s)[index]
	*s = append([]T{}, (*s)[:index]...)
	return item
}
