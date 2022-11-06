package singly

type SLNode[T comparable] struct {
	Value T
	next  *SLNode[T]
	list  *SLList[T]
}

type SLList[T comparable] struct {
	head, tail *SLNode[T]
	len        int
}

func NewSLList[T comparable]() *SLList[T] {
	return new(SLList[T])
}

func (l *SLList[T]) Len() int {
	return l.len
}

func (l *SLList[T]) IsEmpty() bool {
	return l.Len() == 0
}

func (l *SLList[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.len = 0
}

func (l *SLList[T]) Head() *SLNode[T] {
	return l.head
}

func (l *SLList[T]) Tail() *SLNode[T] {
	return l.tail
}

func (l *SLList[T]) ToSlice() []T {
	if l.IsEmpty() {
		return []T{}
	}

	result := []T{}
	node := l.head
	for i := 0; i < l.Len(); i++ {
		result = append(result, node.Value)
		node = node.next
	}
	return result
}

func (l *SLList[T]) InsertHead(newNode *SLNode[T]) *SLNode[T] {
	newNode.next = l.Head()
	newNode.list = l
	l.head = newNode

	if l.Len() == 0 {
		l.tail = newNode
	}

	l.len++
	return newNode
}

func (l *SLList[T]) InsertTail(newNode *SLNode[T]) *SLNode[T] {
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

func (l *SLList[T]) InsertBefore(target, newNode *SLNode[T]) *SLNode[T] {
	if target.list != l {
		return nil
	}

	node := l.Head()
	for i := 0; i < l.Len(); i++ {
		isHeadNodeEqual := i == 0 && node == target
		isNextNodeEqual := node.next != nil && node.next == target

		if isHeadNodeEqual {
			newNode.next = l.Head()
			newNode.list = l
			l.head = newNode
			break
		} else if isNextNodeEqual {
			newNode.next = node.next
			newNode.list = l
			node.next = newNode
			break
		} else if node.next == nil {
			return nil
		}
		node = node.next
	}
	l.len++
	return newNode
}

func (l *SLList[T]) InsertAfter(target, newNode *SLNode[T]) *SLNode[T] {
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

func (l *SLList[T]) DeleteHead() *SLNode[T] {
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

func (l *SLList[T]) DeleteTail() *SLNode[T] {
	if l.IsEmpty() {
		return nil
	}

	node := l.head
	for i := 0; i < l.Len(); i++ {
		if l.isOnlyOneNode() {
			l.head.list = nil
			l.head = nil
			l.tail = nil
			break
		} else if node.next == l.tail {
			dt := l.tail
			l.tail.list = nil
			l.tail = node
			l.tail.next = nil
			node = dt
			break
		}
		node = node.next
	}

	l.len--
	return node
}

func (l *SLList[T]) Delete(target *SLNode[T]) *SLNode[T] {
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

			target.next = nil
			target.list = nil
			l.len--
			break
		}
		node = node.next
	}

	return target
}

func (l *SLList[T]) isOnlyOneNode() bool {
	return l.len == 1
}
