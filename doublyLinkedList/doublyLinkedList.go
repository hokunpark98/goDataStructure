package main

type DNode struct {
	Value int
	Prev  *DNode
	Next  *DNode
}

type DoublyLinkedList struct {
	Head *DNode
	Tail *DNode
	Size int
}

func (l *DoublyLinkedList) Append(value int) {
	newNode := &DNode{Value: value}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		l.Size = 1
		return
	}
	newNode.Prev = l.Tail
	l.Tail.Next = newNode
	l.Tail = newNode
	l.Size++
}

func (l *DoublyLinkedList) Prepend(value int) {
	newNode := &DNode{Value: value}
	if l.Head == nil {
		l.Append(value)
		return
	}
	newNode.Next = l.Head
	l.Head.Prev = newNode
	l.Head = newNode
	l.Size++
}

func (l *DoublyLinkedList) Remove(node *DNode) {
	if node == nil || l.Head == nil {
		return
	}
	l.Size--

	if node == l.Head {
		l.Head = node.Next
		l.Head.Prev = nil
		return
	}

	if node == l.Tail {
		l.Tail = node.Prev
		l.Tail.Next = nil
		return
	}

	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}
