package main

import (
	"container/list"
	"fmt"

	"github.com/denstepanov/algostructs/algorithms"
	"github.com/denstepanov/algostructs/structures/linear/queues"
	"github.com/denstepanov/algostructs/structures/linear/stacks"
)

func main() {
	execStacks()
	execQueues()

	execReverseString()
	list.New()
}

func execStacks() {
	fmt.Println("this is stacks time!")
	fmt.Println("\nStackViaSlice!")
	var stackViaSlice stacks.StackViaSlice[string]

	data := []string{"this", "is", "Sparta"}
	for _, v := range data {
		stackViaSlice.Push(v)
	}

	for len(stackViaSlice) > 0 {
		x, err := stackViaSlice.Pop()
		if err == nil {
			fmt.Println(x)
		}
	}
}

func execQueues() {
	fmt.Println("\nThis is queue time!")
	execQueueViaSlice()
	execQueueViaStack()
}

func execQueueViaSlice() {
	fmt.Println("\nQueueViaSlice")
	var queueViaSlice queues.SliceQueue[int]
	data := []int{3, 541, 300}
	for _, v := range data {
		queueViaSlice.Enqueue(v)
	}

	i, _ := queueViaSlice.Peek()
	fmt.Println(i)

	for queueViaSlice.Len() > 0 {
		item, _ := queueViaSlice.Dequeue()
		fmt.Println(item)
	}

	fmt.Printf("Is queueViaSlice empty? %t!", queueViaSlice.IsEmpty())
}

func execQueueViaStack() {
	fmt.Println("\nQueueViaStack")
	var queueViaStacks queues.StacksQueue[int]
	data := []int{117, 7234, 56}
	for _, v := range data {
		queueViaStacks.Enqueue(v)
	}

	i, _ := queueViaStacks.Peek()
	fmt.Println(i)

	for queueViaStacks.Len() > 0 {
		item, _ := queueViaStacks.Dequeue()
		fmt.Println(item)
	}

	fmt.Printf("Is queueViaStacks empty? %t!", queueViaStacks.IsEmpty())
}

func execReverseString() {
	word := "Hello, World!"
	fmt.Printf("\n Original string: %s", word)
	word, err := algorithms.ReverseString(word)
	if err == nil {
		fmt.Printf("\n Reversed string: %s", word)
	}
}
