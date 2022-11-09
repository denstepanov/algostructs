package slice

type SliceStack[T comparable] []T

func (s *SliceStack[T]) Len() int {
	return len(*s)
}

func (s *SliceStack[T]) IsEmpty() bool {
	return s.Len() == 0
}

func (s *SliceStack[T]) Push(item T) {
	*s = append(*s, item)
}

func (s *SliceStack[T]) Pop() T {
	var result T
	if !s.IsEmpty() {
		index := len(*s) - 1
		result = (*s)[index]
		*s = append([]T{}, (*s)[:index]...)
	}
	return result
}
