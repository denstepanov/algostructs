package list

import "github.com/denstepanov/algostructs/structures/linear/lists/doubly"

type ListStack[T comparable] struct {
	list doubly.List[T]
}

func (s *ListStack[T]) IsEmpty() bool {
	return s.list.IsEmpty()
}

func (s *ListStack[T]) Len() int {
	return s.list.Len()
}

func (s *ListStack[T]) Push(item T) {
	node := &doubly.Node[T]{
		Value: item,
	}
	s.list.InsertTail(node)
}

func (s *ListStack[T]) Pop() T {
	var result T
	if !s.list.IsEmpty() {
		result = s.list.DeleteTail().Value
	}
	return result
}
