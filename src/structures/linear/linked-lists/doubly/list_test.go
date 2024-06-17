package doubly_test

import (
	"fmt"
	"github.com/denstepanov/algostructs/src/utils"
	"testing"

	"github.com/denstepanov/algostructs/src/structures/linear/linked-lists/doubly"
)

// TODO: Детальней проработать сценарии:
// - вставки элементов в середину списка
// - удаления элементов из середины списка
var list = prepare()

func prepare() *doubly.List[string] {
	newList := doubly.New[string]()
	data := []doubly.Node[string]{
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
		t.Fatal(fmt.Printf("DLList.FindByIndex() %s", utils.MethodDoesNotWork))
	}
}

func TestFindByValue(t *testing.T) {
	nodes := list.FindByValue("Mars")

	if nodes[0] != list.Tail() {
		t.Fatal(fmt.Printf("DLList.FindByValue() %s", utils.MethodDoesNotWork))
	}
}

func TestInsertTail(t *testing.T) {
	if list.Len() == 0 {
		t.Fatal(fmt.Printf("DLList.InsertTail() %s", utils.MethodDoesNotWork))
	}
}

func TestClear(t *testing.T) {
	list.Clear()

	if !list.IsEmpty() {
		t.Fatal(fmt.Printf("DLList.Clear() %s", utils.MethodDoesNotWork))
	}

}

func TestToSlice(t *testing.T) {
	listAsSlice := list.ToSlice()

	if len(listAsSlice) != list.Len() {
		t.Fatal(fmt.Printf("DLList.ToSlice() %s", utils.MethodDoesNotWork))
	}
}

func TestInsertHead(t *testing.T) {
	oldLen := list.Len()
	value := "TestInsertHead"
	newNode := &doubly.Node[string]{Value: value}
	newNode = list.InsertHead(newNode)

	if oldLen == list.Len() || newNode == nil {
		t.Fatal("DLList.InsertHead(): new node isn't inserted")
	}
	if newNode.Value != value {
		t.Fatal("DLList.InsertHead(): value doesn't match")
	}
}

func TestInsertBeforeHead(t *testing.T) {
	newNode := &doubly.Node[string]{Value: "TestInsertBeforeHead"}
	newNode = list.InsertBefore(list.Head(), newNode)

	if newNode != list.Head() {
		t.Fatal("DLList.InsertBefore(): new head didn't inserted.")
	}
}

func TestInsertBefore(t *testing.T) {
	oldLen := list.Len()
	newNode := &doubly.Node[string]{Value: "TestInsertBefore"}
	newNode = list.InsertBefore(list.Tail(), newNode)

	if list.Len() == oldLen {
		t.Fatal(fmt.Printf("DLList.InsertBefore() %s", utils.MethodDoesNotWork))
	}
}

func TestInsertAfter(t *testing.T) {
	oldLen := list.Len()
	newNode := &doubly.Node[string]{Value: "TestInsertAfter"}
	newNode = list.InsertAfter(list.Tail(), newNode)

	if list.Len() == oldLen {
		t.Fatal(fmt.Printf("DLList.InsertAfter() %s", utils.MethodDoesNotWork))
	}
}

func TestInsertAfterTail(t *testing.T) {
	newNode := &doubly.Node[string]{Value: "TestInsertAfter"}
	newNode = list.InsertAfter(list.Tail(), newNode)

	if list.Tail() != newNode {
		t.Fatal("DLList.InsertAfter(): new tail didn't inserted.")
	}
}

func TestDeleteHead(t *testing.T) {
	deletedHead := list.DeleteHead()

	if list.Head() == deletedHead {
		t.Fatal("DLList.InsertAfter(): head didn't deleted.")
	}
}

func TestDeleteTail(t *testing.T) {
	deletedTail := list.DeleteTail()

	if list.Tail() == deletedTail {
		t.Fatal("DLList.InsertAfter(): tail didn't deleted.")
	}
}

func TestDelete(t *testing.T) {
}
