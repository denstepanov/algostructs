package slices_test

import (
	"fmt"
	"github.com/denstepanov/algostructs/src/util"
	"testing"

	"github.com/denstepanov/algostructs/src/structures/linear/queues/slices"
)

var queue = prepare()

func prepare() *slices.Queue[int] {
	queue := slices.New[int]()
	data := []int{3, 541, 300}
	for _, v := range data {
		queue.Enqueue(v)
	}
	return queue
}

func TestEnqueue(t *testing.T) {
	if queue.IsEmpty() {
		t.Fatal(fmt.Printf("SliceQueue.Enqueue() %s", util.MethodDoesNotWork))
	}
}

func TestDequeue(t *testing.T) {
	oldLen := queue.Len()
	queue.Dequeue()

	if queue.Len() == oldLen {
		t.Fatal(fmt.Printf("SliceQueue.Dequeue() %s", util.MethodDoesNotWork))
	}
}

func TestPeek(t *testing.T) {
	oldLen := queue.Len()
	queue.Peek()

	if queue.Len() != oldLen {
		t.Fatal(fmt.Printf("SliceQueue.Peek() %s", util.MethodDoesNotWork))
	}
}
