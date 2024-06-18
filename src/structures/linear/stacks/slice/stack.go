package slice

type Stack[T comparable] []T

func (s *Stack[T]) Len() int {
	return len(*s)
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Stack[T]) Push(item T) {
	*s = append(*s, item)
}

func (s *Stack[T]) Pop() T {
	var result T
	if !s.IsEmpty() {
		index := len(*s) - 1
		result = (*s)[index]
		*s = append([]T{}, (*s)[:index]...)
	}
	return result
}
