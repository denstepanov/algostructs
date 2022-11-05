package lists

// TODO: Протестить, устранить недостатки ниже.
// привязка ноды к листу позволяет сразу понять, принадлежит ли эта нода к этому листу или нет.
// вместо ошибок всё же лучше будет возвращать элемент или nil.
// Во втором случае сразу станет понятно что что-то не так

// Во всех for циклах видимо нужно установить l.size - 1?

// как работать со списком, если в нодах могут храниться одинаковые значения?
// как извлекать ноду со значением, которое может дублироваться в рамках всего списка?

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

// TODO: Доработать все методы удаления
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
	isOnlyOneNode := l.head == l.tail
	for i := 0; i < l.Len(); i++ {
		if isOnlyOneNode {
			node.list = nil
			l.head = nil
			l.tail = nil
			break
		} else if node.next == l.tail {
			l.tail.list = nil
			// каким-то образом сохранить значение удалённого хвоста и передать его в качестве результата.
			l.tail = node
			l.tail.next = nil
			break
		}
		node = node.next
	}

	l.len--
	// Подумать, как можно взаимодействовать
	return node
}

func (l *SimplyLinkedList) Delete(target *SimplyLinkedNode) *SimplyLinkedNode {
	if target.list != l {
		return nil
	}

	node := l.head
	for i := 0; i < l.Len(); i++ {
		if target == node {
			l.head = node.next
			target.next = nil
			target.list = l
			break
		} else if target == node.next {
			node.next = node.next.next
			target.next = nil
			target.list = l
			break
		}
		node = node.next
	}
	l.len--
	return target
}
