package list_test

import (
	"fmt"
	"testing"

	"github.com/denstepanov/algostructs/structures"
	"github.com/denstepanov/algostructs/structures/linear/stacks/list"
)

var stack = prepare()

func prepare() *list.ListStack[string] {
	var listStack list.ListStack[string]

	data := []string{"this", "is", "Sparta"}
	for _, v := range data {
		listStack.Push(v)
	}
	return &listStack
}

func TestLen(t *testing.T) {
	if stack.Len() == 0 {
		t.Fatal(fmt.Printf("ListStack.Len() %s", structures.MethodNotWorking))
	}
}

func TestPush(t *testing.T) {
	oldLen := stack.Len()
	stack.Push("Another string")

	if stack.Len() == oldLen {
		t.Fatal(fmt.Printf("ListStack.Push() %s", structures.MethodNotWorking))
	}
}

func TestPop(t *testing.T) {
	oldLen := stack.Len()
	stack.Pop()

	if stack.Len() == oldLen {
		t.Fatal(fmt.Printf("ListStack.Pop() %s", structures.MethodNotWorking))
	}
}
