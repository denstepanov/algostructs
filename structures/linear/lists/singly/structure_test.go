package singly_test

import (
	"fmt"
	"testing"

	"github.com/denstepanov/algostructs/structures"
	"github.com/denstepanov/algostructs/structures/linear/lists/singly"
)

// TODO: Детальней проработать сценарии:
// - вставки элементов в середину списка
// - удаления элементов из середины списка
var list = prepare()

func prepare() *singly.List[string] {
	newList := singly.New[string]()
	data := []singly.Node[string]{
		{
			Value: "Hello",
		},
		{
			Value: "World",
		},
		{
			Value: "Bye",
		},
		{
			Value: "Mars",
		},
	}

	for i := 0; i < len(data); i++ {
		newList.InsertTail(&data[i])
	}

	return newList
}

func TestFindByIndex(t *testing.T) {
	node := list.FindByIndex(3)

	if node != list.Tail() {
		t.Fatal(fmt.Printf("SLList.FindByIndex() %s", structures.MethodNotWorking))
	}
}

func TestFindByValue(t *testing.T) {
	nodes := list.FindByValue("Mars")

	if nodes[0] != list.Tail() {
		t.Fatal(fmt.Printf("SLList.FindByValue() %s", structures.MethodNotWorking))
	}
}

func TestInsertTail(t *testing.T) {
	if list.Len() == 0 {
		t.Fatal(fmt.Printf("SLList.InsertTail() %s", structures.MethodNotWorking))
	}
}

func TestClear(t *testing.T) {
	list.Clear()

	if !list.IsEmpty() {
		t.Fatal(fmt.Printf("SLList.Clear() %s", structures.MethodNotWorking))
	}

}

func TestToSlice(t *testing.T) {
	listAsSlice := list.ToSlice()

	if len(listAsSlice) != list.Len() {
		t.Fatal(fmt.Printf("SLList.ToSlice() %s", structures.MethodNotWorking))
	}
}

func TestInsertHead(t *testing.T) {
	oldLen := list.Len()
	value := "TestInsertHead"
	newNode := &singly.Node[string]{Value: value}
	newNode = list.InsertHead(newNode)

	if oldLen == list.Len() || newNode == nil {
		t.Fatal("SLList.InsertHead(): new node isn't inserted")
	}
	if newNode.Value != value {
		t.Fatal("SLList.InsertHead(): value doesn't match")
	}
}

func TestInsertBeforeHead(t *testing.T) {
	newNode := &singly.Node[string]{Value: "TestInsertBeforeHead"}
	newNode = list.InsertBefore(list.Head(), newNode)

	if newNode != list.Head() {
		t.Fatal("SLList.InsertBefore(): new head didn't inserted.")
	}
}

func TestInsertBefore(t *testing.T) {
	oldLen := list.Len()
	newNode := &singly.Node[string]{Value: "TestInsertBefore"}
	newNode = list.InsertBefore(list.Tail(), newNode)

	if list.Len() == oldLen {
		t.Fatal(fmt.Printf("SLList.InsertBefore() %s", structures.MethodNotWorking))
	}
}

func TestInsertAfter(t *testing.T) {
	oldLen := list.Len()
	newNode := &singly.Node[string]{Value: "TestInsertAfter"}
	newNode = list.InsertAfter(list.Tail(), newNode)

	if list.Len() == oldLen {
		t.Fatal(fmt.Printf("SLList.InsertAfter() %s", structures.MethodNotWorking))
	}
}

func TestInsertAfterTail(t *testing.T) {
	newNode := &singly.Node[string]{Value: "TestInsertAfter"}
	newNode = list.InsertAfter(list.Tail(), newNode)

	if list.Tail() != newNode {
		t.Fatal("SLList.InsertAfter(): new tail didn't inserted.")
	}
}

func TestDeleteHead(t *testing.T) {
	deletedHead := list.DeleteHead()

	if list.Head() == deletedHead {
		t.Fatal("SLList.InsertAfter(): head didn't deleted.")
	}
}

func TestDeleteTail(t *testing.T) {
	deletedTail := list.DeleteTail()

	if list.Tail() == deletedTail {
		t.Fatal("SLList.InsertAfter(): tail didn't deleted.")
	}
}

func TestDelete(t *testing.T) {
}

func TestReverse(t *testing.T) {
	oldHead := list.Head()
	oldTail := list.Tail()
	list.Reverse()

	if list.Head() != oldTail && list.Tail() != oldHead {
		t.Fatal(fmt.Printf("SLList.Reverse() %s", structures.MethodNotWorking))
	}
}
