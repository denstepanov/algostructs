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

func (l *SLList[T]) Head() *SLNode[T] {
	return l.head
}

func (l *SLList[T]) Tail() *SLNode[T] {
	return l.tail
}

func (l *SLList[T]) FindByIndex(idx int) *SLNode[T] {
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

func (l *SLList[T]) FindByValue(value T) []*SLNode[T] {
	result := []*SLNode[T]{}
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

func (l *SLList[T]) ToSlice() []T {
	result := []T{}
	if !l.IsEmpty() {
		node := l.head
		for i := 0; i < l.Len(); i++ {
			result = append(result, node.Value)
			node = node.next
		}
	}

	return result
}

func (l *SLList[T]) isOnlyOneNode() bool {
	return l.Len() == 1
}
