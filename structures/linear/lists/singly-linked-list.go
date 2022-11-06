package lists

// Во всех for циклах видимо нужно установить l.size - 1?

type SimplyLinkedNode struct {
	Value any
	next  *SimplyLinkedNode
	list  *SimplyLinkedList
}

type SimplyLinkedList struct {
	head, tail *SimplyLinkedNode
	len        int
}

func New() *SimplyLinkedList {
	return new(SimplyLinkedList)
}

func (l *SimplyLinkedList) Len() int {
	return l.len
}

func (l *SimplyLinkedList) IsEmpty() bool {
	return l.Len() == 0
}

func (l *SimplyLinkedList) Clear() {
	// А если всё же висящие указатели не будут удаляться?
	l.head = nil
	l.tail = nil
	l.len = 0
}

func (l *SimplyLinkedList) Head() *SimplyLinkedNode {
	return l.head
}

func (l *SimplyLinkedList) Tail() *SimplyLinkedNode {
	return l.tail
}

// реализовать фунукцию вывода всего списка

// функция поиска определённой ноды

// Функция вставки значения внутрь определённой ноды

func (l *SimplyLinkedList) InsertHead(newNode *SimplyLinkedNode) *SimplyLinkedNode {
	newNode.next = l.Head()
	newNode.list = l
	l.head = newNode

	if l.Len() == 0 {
		l.tail = newNode
	}

	l.len++
	return newNode
}

func (l *SimplyLinkedList) InsertTail(newNode *SimplyLinkedNode) *SimplyLinkedNode {
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

func (l *SimplyLinkedList) InsertBefore(target, newNode *SimplyLinkedNode) *SimplyLinkedNode {
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

func (l *SimplyLinkedList) InsertAfter(target, newNode *SimplyLinkedNode) *SimplyLinkedNode {
	if target.list != l {
		return nil
	}

	newNode.next = target.next
	newNode.list = l
	target.next = newNode
	l.len++

	return newNode
}

func (l *SimplyLinkedList) DeleteHead() *SimplyLinkedNode {
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

func (l *SimplyLinkedList) DeleteTail() *SimplyLinkedNode {
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

func (l *SimplyLinkedList) Delete(target *SimplyLinkedNode) *SimplyLinkedNode {
	if target.list != l {
		return nil
	}

	node := l.head
	for i := 0; i < l.Len(); i++ {
		if l.isOnlyOneNode() {
			target.list = nil
			l.head = nil
			l.tail = nil
			break
		} else if target == node.next {
			node.next = node.next.next
			target.next = nil
			target.list = nil
			break
		}
		node = node.next
	}
	l.len--
	return target
}

func (l *SimplyLinkedList) isOnlyOneNode() bool {
	return l.len == 1
}
