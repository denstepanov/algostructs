package linked_list_test

import (
	"fmt"
	"github.com/denstepanov/algostructs/src/utils"
	"testing"

	"github.com/denstepanov/algostructs/src/structures/linear/stacks/linked-list"
)

var stack = prepare()

func prepare() *linked_list.Stack[string] {
	var listStack linked_list.Stack[string]

	data := []string{"this", "is", "Sparta"}
	for _, v := range data {
		listStack.Push(v)
	}
	return &listStack
}

func TestLen(t *testing.T) {
	if stack.Len() == 0 {
		t.Fatal(fmt.Printf("Stack.Len() %s", utils.MethodDoesNotWork))
	}
}

func TestPush(t *testing.T) {
	oldLen := stack.Len()
	stack.Push("Another string")

	if stack.Len() == oldLen {
		t.Fatal(fmt.Printf("Stack.Push() %s", utils.MethodDoesNotWork))
	}
}

func TestPop(t *testing.T) {
	oldLen := stack.Len()
	stack.Pop()

	if stack.Len() == oldLen {
		t.Fatal(fmt.Printf("Stack.Pop() %s", utils.MethodDoesNotWork))
	}
}
