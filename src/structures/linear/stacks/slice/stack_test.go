package slice_test

import (
	"fmt"
	"github.com/denstepanov/algostructs/src/utils"
	"testing"

	"github.com/denstepanov/algostructs/src/structures/linear/stacks/slice"
)

var sliceStack = prepare()

func prepare() *slice.SliceStack[string] {
	var stack slice.SliceStack[string]

	data := []string{"this", "is", "Sparta"}
	for _, v := range data {
		stack.Push(v)
	}
	return &stack
}

func TestLen(t *testing.T) {
	if sliceStack.Len() == 0 {
		t.Fatal(fmt.Printf("SliceStack.Len() %s", utils.MethodNotWorking))
	}
}

func TestPush(t *testing.T) {
	oldLen := sliceStack.Len()
	sliceStack.Push("Another string")

	if sliceStack.Len() == oldLen {
		t.Fatal(fmt.Printf("SliceStack.Push() %s", utils.MethodNotWorking))
	}
}

func TestPop(t *testing.T) {
	oldLen := sliceStack.Len()
	sliceStack.Pop()

	if sliceStack.Len() == oldLen {
		t.Fatal(fmt.Printf("SliceStack.Pop() %s", utils.MethodNotWorking))
	}
}
