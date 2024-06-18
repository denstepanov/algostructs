package stacks_test

import (
	"fmt"
	"github.com/denstepanov/algostructs/src/utils"
	"testing"

	"github.com/denstepanov/algostructs/src/structures/linear/queues/stacks"
)

var queue = prepare()

func prepare() *stacks.Queue[int] {
	queue := stacks.New[int]()
	data := []int{3, 541, 300}
	for _, v := range data {
		queue.Enqueue(v)
	}
	return queue
}

func TestEnqueue(t *testing.T) {
	if queue.IsEmpty() {
		t.Fatal(fmt.Printf("StacksQueue.Enqueue() %s", utils.MethodDoesNotWork))
	}
}

func TestDequeue(t *testing.T) {
	oldLen := queue.Len()
	queue.Dequeue()

	if queue.Len() == oldLen {
		t.Fatal(fmt.Printf("StacksQueue.Dequeue() %s", utils.MethodDoesNotWork))
	}
}

func TestPeek(t *testing.T) {
	oldLen := queue.Len()
	queue.Peek()

	if queue.Len() != oldLen {
		t.Fatal(fmt.Printf("StacksQueue.Peek() %s", utils.MethodDoesNotWork))
	}
}
