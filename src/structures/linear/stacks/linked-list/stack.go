package linked_list

import "github.com/denstepanov/algostructs/src/structures/linear/linked-lists/doubly"

type Stack[T comparable] struct {
	list doubly.LinkedList[T]
}

func (s *Stack[T]) IsEmpty() bool {
	return s.list.IsEmpty()
}

func (s *Stack[T]) Len() int {
	return s.list.Len()
}

func (s *Stack[T]) Push(item T) {
	node := &doubly.Node[T]{
		Value: item,
	}
	s.list.InsertTail(node)
}

func (s *Stack[T]) Pop() T {
	var result T
	if !s.list.IsEmpty() {
		result = s.list.DeleteTail().Value
	}
	return result
}
