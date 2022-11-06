package main

import (
	"fmt"

	"github.com/denstepanov/algostructs/algorithms"
	"github.com/denstepanov/algostructs/structures/linear/lists"
	"github.com/denstepanov/algostructs/structures/linear/queues"
	"github.com/denstepanov/algostructs/structures/linear/stacks"
)

func main() {
	execStacks()
	execQueues()
	execLists()

	execReverseString()
}

func execStacks() {
	fmt.Println("this is stacks time!")
	fmt.Println("\nSliceStack!")
	var sliceStack stacks.SliceStack[string]

	data := []string{"this", "is", "Sparta"}
	for _, v := range data {
		sliceStack.Push(v)
	}

	for len(sliceStack) > 0 {
		x, err := sliceStack.Pop()
		if err == nil {
			fmt.Println(x)
		}
	}
}

func execLists() {
	fmt.Println("\nDat Lists time!")
	execsll()
}

func execsll() {
	fmt.Println("This is Simply Linked List")
	sll := lists.NewSimplyLinkedList[string]()

	fmt.Printf("Pointer to the sll: %p\n", sll)
	head := &lists.SimplyLinkedNode[string]{Value: "Hello, World!"}

	sll.InsertHead(head)
	fmt.Printf("Print head struct: %v\n", sll.Head())

	tail := &lists.SimplyLinkedNode[string]{Value: "World, Hello!"}

	sll.InsertTail(tail)
	fmt.Printf("List length: %d\n", sll.Len())
	fmt.Printf("Print tail struct: %v\n", sll.Tail())

	// sll.Clear()
	// fmt.Printf("List length after clear method: %d\n", sll.Len())

	// fmt.Printf("Old tail struct: %v\n", sll.Tail())
	// sll.Delete(sll.Tail())
	// fmt.Printf("New tail struct: %v\n", sll.Tail())

	// fmt.Printf("Old head struct: %v\n", sll.Head())
	// sll.Delete(sll.Head())
	// fmt.Printf("New head struct: %v\n", sll.Head())

	// sll.DeleteHead()
	// fmt.Printf("Print head struct after delete previous head: %v\n", sll.Head())

	// sll.DeleteTail()
	// fmt.Printf("Print head struct after delete previous tail: %v\n", sll.Tail())

	newNode := &lists.SimplyLinkedNode[string]{Value: "Bb, Mars!!!"}

	// sll.InsertBefore(head, newNode)
	// fmt.Printf("New head: %v\n", sll.Head())
	// fmt.Println(sll.Len())

	sll.InsertAfter(tail, newNode)
	fmt.Printf("New tail after InsertAfter: %v\n", sll.Tail())
	fmt.Println(sll.Len())

	fmt.Printf("Len before ToSlice: %d", sll.Len())
	fmt.Println(sll.ToSlice())
}

func execQueues() {
	fmt.Println("\nThis is queue time!")
	execSliceQueue()
	execStacksQueue()
}

func execSliceQueue() {
	fmt.Println("\nSliceQueue!")
	var sliceQueue queues.SliceQueue[int]
	data := []int{3, 541, 300}
	for _, v := range data {
		sliceQueue.Enqueue(v)
	}

	i, _ := sliceQueue.Peek()
	fmt.Println(i)

	for sliceQueue.Len() > 0 {
		item, _ := sliceQueue.Dequeue()
		fmt.Println(item)
	}

	fmt.Printf("Is sliceQueue empty? %t!", sliceQueue.IsEmpty())
}

func execStacksQueue() {
	fmt.Println("\nQueueViaStack")
	var stacksQueue queues.StacksQueue[int]
	data := []int{117, 7234, 56}
	for _, v := range data {
		stacksQueue.Enqueue(v)
	}

	i, _ := stacksQueue.Peek()
	fmt.Println(i)

	for stacksQueue.Len() > 0 {
		item, _ := stacksQueue.Dequeue()
		fmt.Println(item)
	}

	fmt.Printf("Is stacksQueue empty? %t!", stacksQueue.IsEmpty())
}

func execReverseString() {
	word := "Hello, World!"
	fmt.Printf("\n Original string: %s", word)
	word, err := algorithms.ReverseString(word)
	if err == nil {
		fmt.Printf("\n Reversed string: %s", word)
	}
}
