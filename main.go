package main

import (
	"fmt"

	"github.com/denstepanov/algostructs/algorithms"
	"github.com/denstepanov/algostructs/structures/linear/queues"
	"github.com/denstepanov/algostructs/structures/linear/stacks"
)

func main() {
	execStacks()
	execQueues()

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
