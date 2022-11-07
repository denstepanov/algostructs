package list_test

import (
	"testing"

	"github.com/denstepanov/algostructs/structures/linear/queues/list"
)

func TestDequeue(t *testing.T) {
	queue := prepare(t)
	oldLen := queue.Len()
	queue.Dequeue()
	if queue.Len() == oldLen {
		t.Fatal("listQueue.Dequeue() method doesn't work.")
	}
}

func TestPeek(t *testing.T) {
	queue := prepare(t)
	oldLen := queue.Len()
	queue.Peek()
	if queue.Len() != oldLen {
		t.Fatal("listQueue.Peek() method doesn't work.")
	}
}

func prepare(t *testing.T) *list.ListQueue[int] {
	queue := list.New[int]()
	data := []int{3, 541, 300}
	for _, v := range data {
		queue.Enqueue(&v)
	}
	if queue.IsEmpty() {
		t.Fatal("listQueue.Enqueue() doesn't work.")
	}

	return queue
}
