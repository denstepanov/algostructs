package list_test

import (
	"fmt"
	"testing"

	"github.com/denstepanov/algostructs/structures"
	"github.com/denstepanov/algostructs/structures/linear/queues/list"
)

var queue = prepare()

func prepare() *list.ListQueue[int] {
	queue := list.New[int]()
	data := []int{3, 541, 300}
	for _, v := range data {
		queue.Enqueue(v)
	}
	return queue
}

func TestEnqueue(t *testing.T) {
	if queue.IsEmpty() {
		t.Fatal(fmt.Printf("ListQueue.Enqueue() %s", structures.MethodNotWorking))
	}
}

func TestDequeue(t *testing.T) {
	oldLen := queue.Len()
	queue.Dequeue()

	if queue.Len() == oldLen {
		t.Fatal(fmt.Printf("ListQueue.Dequeue() %s", structures.MethodNotWorking))
	}
}

func TestPeek(t *testing.T) {
	oldLen := queue.Len()
	queue.Peek()

	if queue.Len() != oldLen {
		t.Fatal(fmt.Printf("ListQueue.Peek() %s", structures.MethodNotWorking))
	}
}
