package stacks

import (
	"errors"

	"github.com/denstepanov/algostructs/structures"
)

type StackViaSlice[T comparable] []T

func (s *StackViaSlice[T]) IsEmpty() bool {
	return len(*s) == 0
}

func (s *StackViaSlice[T]) Len() int {
	return len(*s)
}

func (s *StackViaSlice[T]) Push(item T) {
	*s = append(*s, item)
}

func (s *StackViaSlice[T]) Pop() (item T, err error) {
	if s.IsEmpty() {
		err = errors.New(structures.EmptyStack)
		return item, err
	}
	index := len(*s) - 1
	item = (*s)[index]
	*s = append([]T{}, (*s)[:index]...)
	return item, nil
}
