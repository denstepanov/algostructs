package stacks_test

import (
	"fmt"
	"testing"

	"github.com/denstepanov/algostructs/structures"
	"github.com/denstepanov/algostructs/structures/linear/queues/stacks"
)

var queue = prepare()

func prepare() *stacks.StacksQueue[int] {
	queue := stacks.New[int]()
	data := []int{3, 541, 300}
	for _, v := range data {
		queue.Enqueue(v)
	}
	return queue
}

func TestEnqueue(t *testing.T) {
	if queue.IsEmpty() {
		t.Fatal(fmt.Printf("StacksQueue.Enqueue() %s", structures.MethodNotWorking))
	}
}

func TestDequeue(t *testing.T) {
	oldLen := queue.Len()
	queue.Dequeue()

	if queue.Len() == oldLen {
		t.Fatal(fmt.Printf("StacksQueue.Dequeue() %s", structures.MethodNotWorking))
	}
}

func TestPeek(t *testing.T) {
	oldLen := queue.Len()
	queue.Peek()

	if queue.Len() != oldLen {
		t.Fatal(fmt.Printf("StacksQueue.Peek() %s", structures.MethodNotWorking))
	}
}
