package main

import (
	"reflect"
	"testing"
)

func TestLinkedList_Append(t *testing.T) {
	list := &LinkedList{}
	list.Append(10)
	list.Append(20)

	// 기대: [10, 20]
	want := []int{10, 20}
	got := toSlice(list) // 아래에 별도 유틸 함수 정의
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Append() = %v, want %v", got, want)
	}
}

func TestLinkedList_Prepend(t *testing.T) {
	list := &LinkedList{}
	list.Prepend(10)
	list.Prepend(20)

	// 기대: [20, 10]
	want := []int{20, 10}
	got := toSlice(list)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Prepend() = %v, want %v", got, want)
	}
}

func TestLinkedList_Insert(t *testing.T) {
	list := &LinkedList{}
	list.Append(10)
	list.Append(20)
	list.Append(30)

	list.Insert(1, 15) // 10, 15, 20, 30
	want := []int{10, 15, 20, 30}
	got := toSlice(list)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Insert() = %v, want %v", got, want)
	}
}

func TestLinkedList_Remove(t *testing.T) {
	list := &LinkedList{}
	list.Append(10)
	list.Append(20)
	list.Append(30)

	list.Remove(1) // 인덱스 1(값=20) 제거 -> 10, 30
	want := []int{10, 30}
	got := toSlice(list)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Remove() = %v, want %v", got, want)
	}
}

func TestLinkedList_Search(t *testing.T) {
	list := &LinkedList{}
	list.Append(10)
	list.Append(20)
	list.Append(30)

	if idx := list.Search(10); idx != 0 {
		t.Errorf("Search(10) = %d, want 0", idx)
	}
	if idx := list.Search(30); idx != 2 {
		t.Errorf("Search(30) = %d, want 2", idx)
	}
	if idx := list.Search(999); idx != -1 {
		t.Errorf("Search(999) = %d, want -1", idx)
	}
}

// toSlice: LinkedList를 []int로 변환하는 유틸 함수 (테스트용)
func toSlice(list *LinkedList) []int {
	var result []int
	cur := list.Head
	for cur != nil {
		result = append(result, cur.Value)
		cur = cur.Next
	}
	return result
}
