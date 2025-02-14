package main

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		index int
		value int
		want  []int
	}{
		{
			name:  "중간 삽입",
			input: []int{1, 2, 3},
			index: 1,
			value: 999,
			want:  []int{1, 999, 2, 3},
		},
		{
			name:  "맨 앞 삽입",
			input: []int{10, 20, 30},
			index: 0,
			value: 5,
			want:  []int{5, 10, 20, 30},
		},
		{
			name:  "맨 뒤 삽입",
			input: []int{100, 200},
			index: 2,
			value: 300,
			want:  []int{100, 200, 300},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Insert(tt.input, tt.index, tt.value)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		index int
		want  []int
	}{
		{
			name:  "중간 삭제",
			input: []int{1, 2, 3, 4},
			index: 2, // 세 번째 요소(3) 삭제
			want:  []int{1, 2, 4},
		},
		{
			name:  "맨 앞 삭제",
			input: []int{10, 20, 30},
			index: 0,
			want:  []int{20, 30},
		},
		{
			name:  "맨 뒤 삭제",
			input: []int{100, 200, 300},
			index: 2,
			want:  []int{100, 200},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Remove(tt.input, tt.index)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}
