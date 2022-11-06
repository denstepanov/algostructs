package stacks

import (
	"errors"

	"github.com/denstepanov/algostructs/structures"
)

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

func (s *SliceStack[T]) Pop() (item T, err error) {
	if s.IsEmpty() {
		err = errors.New(structures.EmptyStack)
		return item, err
	}
	index := len(*s) - 1
	item = (*s)[index]
	*s = append([]T{}, (*s)[:index]...)
	return item, nil
}
