package singly

type Node[T comparable] struct {
	Value T
	next  *Node[T]
	list  *LinkedList[T]
}

func (n *Node[T]) clearNode() {
	n.list = nil
	n.next = nil
}

type LinkedList[T comparable] struct {
	head, tail *Node[T]
	len        int
}

func New[T comparable]() *LinkedList[T] {
	return new(LinkedList[T])
}

func (l *LinkedList[T]) Len() int {
	return l.len
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.Len() == 0
}

func (l *LinkedList[T]) Clear() {
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

func (l *LinkedList[T]) Head() *Node[T] {
	return l.head
}

func (l *LinkedList[T]) Tail() *Node[T] {
	return l.tail
}

func (l *LinkedList[T]) FindByIndex(idx int) *Node[T] {
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

func (l *LinkedList[T]) FindByValue(value T) []*Node[T] {
	var result []*Node[T]
	if !l.IsEmpty() {
		node := l.head
		for i := 0; i <= l.len-1; i++ {
			if node.Value == value {
				result = append(result, node)
			}
			node = node.next
		}
	}
	return result
}

func (l *LinkedList[T]) InsertHead(newNode *Node[T]) *Node[T] {
	newNode.next = l.Head()
	newNode.list = l
	l.head = newNode

	if l.Len() == 0 {
		l.tail = newNode
	}

	l.len++
	return newNode
}

func (l *LinkedList[T]) InsertTail(newNode *Node[T]) *Node[T] {
	newNode.list = l

	if l.IsEmpty() {
		l.head = newNode
	} else {
		l.tail.next = newNode
	}

	l.tail = newNode
	l.len++
	return newNode
}

func (l *LinkedList[T]) InsertBefore(target, newNode *Node[T]) *Node[T] {
	if target.list != l {
		return nil
	}

	node := l.Head()
	for i := 0; i < l.Len(); i++ {
		isHeadNodeEqual := l.head == target
		isNextNodeEqual := node.next == target

		if isHeadNodeEqual {
			l.InsertHead(newNode)
			break
		} else if isNextNodeEqual {
			newNode.next = node.next
			newNode.list = l
			node.next = newNode
			l.len++
			break
		}
		node = node.next
	}
	return newNode
}

func (l *LinkedList[T]) InsertAfter(target, newNode *Node[T]) *Node[T] {
	if target.list != l {
		return nil
	}

	newNode.next = target.next
	newNode.list = l
	target.next = newNode
	if newNode.next == nil {
		l.tail = newNode
	}
	l.len++

	return newNode
}

func (l *LinkedList[T]) DeleteHead() *Node[T] {
	if l.IsEmpty() {
		return nil
	}

	node := l.head
	node.list = nil
	if node == l.tail {
		l.head = nil
		l.tail = nil
	}

	if node.next != nil {
		l.head = node.next
		node.next = nil
	}

	l.len--
	return node
}

func (l *LinkedList[T]) DeleteTail() *Node[T] {
	if l.IsEmpty() {
		return nil
	}

	node := l.head
	for i := 0; i < l.Len(); i++ {
		if l.isOnlyOneNode() {
			l.Clear()
			break
		} else if node.next == l.tail {
			dt := l.tail
			l.tail.list = nil
			l.tail = node
			l.tail.next = nil
			node = dt
			l.len--
			break
		}
		node = node.next
	}
	return node
}

func (l *LinkedList[T]) Delete(target *Node[T]) *Node[T] {
	if target.list != l {
		return nil
	}

	if l.isOnlyOneNode() {
		l.Clear()
		return target
	}

	if target == l.head {
		l.head = l.head.next
		target.next = nil
		l.len--
		return target
	}

	node := l.head
	for i := 0; i < l.Len(); i++ {
		if target == node.next {
			if node.next == l.tail {
				l.tail = node
			}
			node.next = target.next
			target.clearNode()
			l.len--
			break
		}
		node = node.next
	}

	return target
}

func (l *LinkedList[T]) Reverse() {
	if !l.IsEmpty() {
		next := l.head.next
		current := l.head
		var prev *Node[T] = nil
		for current != nil {
			current.next = prev
			if prev == nil {
				l.tail = current
			}
			current = next
			if next != nil {
				next = next.next
			}
			prev = current
		}
		l.head = prev
	}
}

func (l *LinkedList[T]) ToSlice() []T {
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

func (l *LinkedList[T]) isOnlyOneNode() bool {
	return l.Len() == 1
}
