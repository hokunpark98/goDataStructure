package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
	Size int
}

// Append: 리스트 뒤쪽에 새 노드를 추가
func (l *LinkedList) Append(value int) {
	newNode := &Node{Value: value, Next: nil}

	if l.Head == nil {
		l.Head = newNode
		l.Size++
		return
	}

	// 마지막 노드를 찾아서 연결
	cur := l.Head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = newNode
	l.Size++
}

// Prepend: 리스트 맨 앞쪽에 새 노드를 추가
func (l *LinkedList) Prepend(value int) {
	newNode := &Node{Value: value, Next: nil}
	newNode.Next = l.Head
	l.Head = newNode
	l.Size++
}

// Insert: index 위치에 새 노드를 삽입
func (l *LinkedList) Insert(index int, value int) {
	if index <= 0 {
		l.Prepend(value)
		return
	}
	if index >= l.Size {
		l.Append(value)
		return
	}

	newNode := &Node{Value: value}
	prev := l.Head
	for i := 0; i < index-1; i++ {
		prev = prev.Next
	}

	newNode.Next = prev.Next
	prev.Next = newNode
	l.Size++
}

// Remove: index 위치의 노드를 삭제
func (l *LinkedList) Remove(index int) {
	if l.Head == nil || index < 0 || index >= l.Size {
		return
	}
	if index == 0 {
		l.Head = l.Head.Next
		l.Size--
	}
	prev := l.Head
	for i := 0; i < index-1; i++ {
		prev = prev.Next
	}
	prev.Next = prev.Next.Next
	l.Size--
}

// Search: value를 가진 노드를 찾아 인덱스를 반환
func (l *LinkedList) Search(value int) int {
	index := 0
	cur := l.Head

	for cur != nil {
		if cur.Value == value {
			return index
		}
		cur = cur.Next
		index++
	}
	return -1
}

func (l *LinkedList) Print() {
	cur := l.Head

	for cur != nil {
		fmt.Printf("%d -> ", cur.Value)
		cur = cur.Next
	}
	fmt.Println("nil")
}
