package slice_test

import (
	"testing"

	"github.com/denstepanov/algostructs/structures/linear/stacks/slice"
)

func TestLen(t *testing.T) {
	sliceStack := prepare()
	if sliceStack.Len() == 0 {
		t.Fatal("sliceStack.Len() method isn't working correctly.")
	}
}

func TestPush(t *testing.T) {
	sliceStack := prepare()
	stackOldLen := sliceStack.Len()
	sliceStack.Push("Another string")
	if sliceStack.Len() == stackOldLen {
		t.Fatal("New string didn't push to the slice stack!")
	}
}

func TestPop(t *testing.T) {
	sliceStack := prepare()
	stackOldLen := sliceStack.Len()
	sliceStack.Pop()
	if sliceStack.Len() == stackOldLen {
		t.Fatal("String didn't pop from slice stack!")
	}
}

func prepare() *slice.SliceStack[string] {
	var sliceStack slice.SliceStack[string]

	data := []string{"this", "is", "Sparta"}
	for _, v := range data {
		sliceStack.Push(v)
	}
	return &sliceStack
}
