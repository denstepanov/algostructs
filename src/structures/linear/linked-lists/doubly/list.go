package doubly

type Node[T comparable] struct {
	next, prev *Node[T]
	Value      T
	list       *List[T]
}

func (n *Node[T]) clearNode() {
	n.list = nil
	n.next = nil
	n.prev = nil
}

type List[T comparable] struct {
	head, tail *Node[T]
	len        int
}

func New[T comparable]() *List[T] {
	return new(List[T])
}

func (l *List[T]) Len() int {
	return l.len
}

func (l *List[T]) IsEmpty() bool {
	return l.Len() == 0
}

func (l *List[T]) Clear() {
	if l.Len() > 0 {
		node := l.head
		for i := 0; i < l.Len(); i++ {
			node.list = nil
			node = node.next
		}

		l.head = nil
		l.tail = nil
		l.len = 0
	}
}

func (l *List[T]) Head() *Node[T] {
	return l.head
}

func (l *List[T]) Tail() *Node[T] {
	return l.tail
}

func (l *List[T]) FindByIndex(idx int) *Node[T] {
	if l.Len()-1 < idx {
		return nil
	}

	node := l.head
	for i := 0; i < idx; i++ {
		if i == idx {
			break
		}
		node = node.next
	}
	return node
}

func (l *List[T]) FindByValue(value T) []*Node[T] {
	result := []*Node[T]{}
	if l.IsEmpty() {
		return result
	}

	node := l.head
	for i := 0; i < l.Len(); i++ {
		if node.Value == value {
			result = append(result, node)
		}
		node = node.next
	}
	return result
}

func (l *List[T]) InsertHead(nn *Node[T]) *Node[T] {
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

func (l *List[T]) InsertTail(nn *Node[T]) *Node[T] {
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

func (l *List[T]) InsertBefore(t, nn *Node[T]) *Node[T] {
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

func (l *List[T]) InsertAfter(t, nn *Node[T]) *Node[T] {
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

func (l *List[T]) DeleteHead() *Node[T] {
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

func (l *List[T]) DeleteTail() *Node[T] {
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

func (l *List[T]) Delete(t *Node[T]) *Node[T] {
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
		t.clearNode()
		l.len--
	} else if t.next == l.Tail() {
		l.tail.prev = t.prev
		l.tail.prev.next = l.tail
		t.clearNode()
		l.len--
	} else {
		// Сценарий удаления элемента где-то в середине списка.
		// Счёт с 3го элемента.
		node := l.head.next
		for i := 0; i < l.Len(); i++ {
			if node.next == t {
				node.next = t.next
				node.next.prev = node
				t.clearNode()
			}
			node = node.next
		}
		l.len--
	}
	return t
}

func (l *List[T]) ToSlice() []T {
	var result []T
	if !l.IsEmpty() {
		node := l.head
		for i := 0; i < l.Len(); i++ {
			result = append(result, node.Value)
			node = node.next
		}
	}

	return result
}
