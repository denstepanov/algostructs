package linked_lists_test

import (
	"fmt"
	"github.com/denstepanov/algostructs/src/util"
	"testing"

	"github.com/denstepanov/algostructs/src/structures/linear/queues/linked-lists"
)

var queue = prepare()

// TODO: Вычленить подготовку в отдельный тест Enqueue. Сделать это для всех видов очередей.
func prepare() *linked_lists.Queue[int] {
	queue := linked_lists.New[int]()
	data := []int{3, 541, 300}
	for _, v := range data {
		queue.Enqueue(v)
	}
	return queue
}

func TestEnqueue(t *testing.T) {
	if queue.IsEmpty() {
		t.Fatal(fmt.Printf("Queue.Enqueue() %s", util.MethodDoesNotWork))
	}
}

func TestDequeue(t *testing.T) {
	oldLen := queue.Len()
	queue.Dequeue()

	if queue.Len() == oldLen {
		t.Fatal(fmt.Printf("Queue.Dequeue() %s", util.MethodDoesNotWork))
	}
}

func TestPeek(t *testing.T) {
	oldLen := queue.Len()
	queue.Peek()

	if queue.Len() != oldLen {
		t.Fatal(fmt.Printf("Queue.Peek() %s", util.MethodDoesNotWork))
	}
}
