package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(value int) {
	s.items = append(s.items, value)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	topIndex := len(s.items) - 1
	value := s.items[topIndex]
	s.items = s.items[:topIndex]
	return value, true
}

func (s *Stack) Peek() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}

	topIndex := len(s.items) - 1
	value := s.items[topIndex]
	return value, true
}

func (s *Stack) IsEmpty() bool {
	if len(s.items) == 0 {
		return true
	} else {
		return false
	}
}

func (s *Stack) Size() int {
	return len(s.items)
}

func main() {
	stack := &Stack{}

	fmt.Println("스택에 10, 20, 30 을 차례대로 Push")
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	fmt.Printf("현재 스택 크기: %d\n", stack.Size())

	// 스택의 맨 위 값 확인
	if top, ok := stack.Peek(); ok {
		fmt.Printf("Peek: 스택의 top 값은 %d 입니다.\n", top)
	}

	// 스택의 모든 요소를 Pop 해보기
	fmt.Println("스택의 요소들을 Pop:")
	for !stack.IsEmpty() {
		if value, ok := stack.Pop(); ok {
			fmt.Println(value)
		}
	}

	fmt.Printf("스택 비었는지 확인: %v\n", stack.IsEmpty())
}
