package doubly

type DLNode[T comparable] struct {
	next, prev *DLNode[T]
	Value      *T
	list       *DLList[T]
}

type DLList[T comparable] struct {
	head, tail *DLNode[T]
	len        int
}

func NewDLList[T comparable]() *DLList[T] {
	return new(DLList[T])
}

func (l *DLList[T]) Len() int {
	return l.len
}

func (l *DLList[T]) IsEmpty() bool {
	return l.Len() == 0
}

func (l *DLList[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.len = 0
}

func (l *DLList[T]) Head() *DLNode[T] {
	return l.head
}

func (l *DLList[T]) Tail() *DLNode[T] {
	return l.tail
}

func (l *DLList[T]) Find(value T) *DLNode[T] {
	var result *DLNode[T]
	node := l.Head()
	for i := 0; i < l.Len(); i++ {
		if node.Value == &value {
			result = node
			break
		}
	}
	return result
}

func (l *DLList[T]) InsertHead(nn *DLNode[T]) *DLNode[T] {
	if l.IsEmpty() {
		l.tail = nn
	} else {
		l.head.prev = nn
		nn.next = l.Head()
	}
	nn.list = l
	l.head = nn
	l.len++
	return nn
}

func (l *DLList[T]) InsertTail(nn *DLNode[T]) *DLNode[T] {
	if l.IsEmpty() {
		l.head = nn
	} else {
		l.tail.next = nn
		nn.prev = l.Tail()
	}
	nn.list = l
	l.tail = nn
	l.len++
	return nn
}

func (l *DLList[T]) InsertBefore(t, nn *DLNode[T]) *DLNode[T] {
	if t.list != l {
		return nil
	}

	if t == l.Head() {
		l.InsertHead(nn)
	} else {
		nn.prev = t.prev
		nn.next = t
		nn.list = l
		t.prev = nn
		l.len++
	}
	return nn
}

func (l *DLList[T]) InsertAfter(t, nn *DLNode[T]) *DLNode[T] {
	if t.list != l {
		return nil
	}

	if t == l.Tail() {
		l.InsertTail(nn)
	} else {
		nn.next = t.next
		nn.prev = t
		nn.list = l
		t.next = nn
		l.len++
	}

	return nn
}

func (l *DLList[T]) DeleteHead() *DLNode[T] {
	if l.IsEmpty() {
		return nil
	}

	node := l.head
	node.list = nil
	if l.Head() == l.Tail() {
		l.Clear()
	} else {
		l.head = l.head.next
		node.next = nil
		l.len--
	}
	return node
}

func (l *DLList[T]) DeleteTail() *DLNode[T] {
	if l.IsEmpty() {
		return nil
	}

	node := l.tail
	node.list = nil
	if l.Head() == l.Tail() {
		l.Clear()
	} else {
		l.tail = l.tail.prev
		node.prev = nil
		l.len--
	}
	return node
}

func (l *DLList[T]) Delete(t *DLNode[T]) *DLNode[T] {
	if t.list != l {
		return nil
	}

	if l.Head() == l.Tail() {
		l.Clear()
	} else if t == l.Head() {
		l.DeleteHead()
	} else if t == l.Tail() {
		l.DeleteTail()
	} else if t.prev == l.Head() {
		l.head.next = t.next
		l.head.next.prev = l.head // обратная ссылка нового элемента, следующего за head.
		clearNode(t)
		l.len--
	} else if t.next == l.Tail() {
		l.tail.prev = t.prev
		l.tail.prev.next = l.tail
		clearNode(t)
		l.len--
	} else {
		// Сценарий удаления элемента где-то в середине списка.
		// Счёт с 3го элемента.
		node := l.head.next
		for i := 0; i < l.Len(); i++ {
			if node.next == t {
				node.next = t.next
				node.next.prev = node
				clearNode(t)
			}
			node = node.next
		}
		l.len--
	}
	return t
}

func clearNode[T comparable](node *DLNode[T]) {
	node.list = nil
	node.next = nil
	node.prev = nil
}
