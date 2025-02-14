package main

import "fmt"

func Insert(slice []int, index, value int) []int {
	if index < 0 || index > len(slice) {
		panic("index out of range")
	}
	// append로 공간 1개 늘리기
	slice = append(slice, 0)

	// 추가할 인덱스 자리부터 뒤에있는애들 뒤로 한 칸씩 밀기
	copy(slice[index+1:], slice[index:])
	slice[index] = value
	return slice
}

func Remove(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		panic("index out of range")
	}
	// 삭제할 인덱스 뒤에있는 애들을 한 칸씩 땡기기
	copy(slice[index:], slice[index+1:])
	slice = slice[:len(slice)-1]
	return slice
}

func main() {
	fmt.Println("====배열/슬라이스 연습====")
	// 1. 배열 출력
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Printf("1. 배열 arr=%v\n", arr)

	// 2. 슬라이스 출력
	var slice []int = []int{1, 2, 3}
	fmt.Printf("2. 슬라이스 초기값 slice=%v, len=%d, cap=%d\n", slice, len(slice), cap(slice))

	// 3. append로 확장
	slice = append(slice, 4)
	fmt.Printf("3. 슬라이스 append slice=%v, len=%d, cap=%d\n", slice, len(slice), cap(slice))

	// 4. 부분 슬라이즈
	sub := slice[1:3]
	fmt.Printf("4. 슬라이스 [1:3]=%v", sub)

	// 5. Insert 함수 사용 (중간 삽입)
	slice = Insert(slice, 2, 42)
	fmt.Printf("5. Insert 중간 slice=%v, len=%d, cap=%d\n", slice, len(slice), cap(slice))

	// 6. Remove 함수 사용 (중간 삭제)
	slice = Remove(slice, 3)
	fmt.Printf("6. Remove 중간 slice=%v, len=%d, cap=%d\n", slice, len(slice), cap(slice))

	// 7. make로 슬라이스 생성
	slice2 := make([]int, 5, 10)
	fmt.Printf("7. make로 생성 slice2 = %v, len = %d, cap = %d\n", slice2, len(slice2), cap(slice2))
}
