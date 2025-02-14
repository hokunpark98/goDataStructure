package main

import "fmt"

func main() {
	list := &LinkedList{}

	fmt.Println("===== 단일 연결 리스트 데모 =====")
	list.Append(10)
	list.Append(20)
	list.Append(30)
	list.Print() // 10 -> 20 -> 30 -> nil

	list.Prepend(5)
	list.Print() // 5 -> 10 -> 20 -> 30 -> nil

	list.Insert(2, 15)
	list.Print() // 5 -> 10 -> 15 -> 20 -> 30 -> nil

	list.Remove(1) // 인덱스 1 (값 10) 제거
	list.Print()   // 5 -> 15 -> 20 -> 30 -> nil

	idx := list.Search(20) // 값 20은 인덱스 2
	fmt.Println("Search(20) =", idx)

	fmt.Println("Size =", list.Size) // 총 노드 수
}
