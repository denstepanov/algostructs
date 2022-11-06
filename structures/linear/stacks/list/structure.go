package list

import "github.com/denstepanov/algostructs/structures/linear/lists/doubly"

type ListStack[T comparable] struct {
	list doubly.DLList[T]
}

func (s *ListStack[T]) IsEmpty() bool {
	return s.list.IsEmpty()
}

func (s *ListStack[T]) Len() int {
	return s.list.Len()
}

func (s *ListStack[T]) Push(item T) {
	node := &doubly.DLNode[T]{
		Value: &item,
	}
	s.list.InsertTail(node)
}

func (s *ListStack[T]) Pop() (item T) {
	if s.list.IsEmpty() {
		return item
	}

	return *s.list.DeleteTail().Value
}
