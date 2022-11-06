package list_test

import (
	"testing"

	"github.com/denstepanov/algostructs/structures/linear/stacks/list"
)

func TestLen(t *testing.T) {
	stack := prepare()
	if stack.Len() == 0 {
		t.Fatal("sliceStack.Len() method isn't working correctly.")
	}
}

func TestPush(t *testing.T) {
	stack := prepare()
	stackOldLen := stack.Len()
	stack.Push("Another string")
	if stack.Len() == stackOldLen {
		t.Fatal("New string didn't push to the slice stack!")
	}
}

func TestPop(t *testing.T) {
	stack := prepare()
	stackOldLen := stack.Len()
	stack.Pop()
	if stack.Len() == stackOldLen {
		t.Fatal("String didn't pop from slice stack!")
	}
}

func prepare() *list.ListStack[string] {
	var stack list.ListStack[string]

	data := []string{"this", "is", "Sparta"}
	for _, v := range data {
		stack.Push(v)
	}
	return &stack
}
