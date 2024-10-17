package main

import (
	"reflect"
	"testing"
)

// TestSliceSlice проверяет работу функции sliceslice
func TestSliceSlice(t *testing.T) {
	tests := []struct {
		matrix   [][]int
		expected []int
	}{
		{
			matrix: [][]int{
				{3, 3, 1},
				{4, 0, 4},
				{3, 1, 1},
			},
			expected: []int{3, 3, 4}, // максимальные значения по строкам: 3, 4, 3 -> отсортировано: [3, 3, 4]
		},
		{
			matrix: [][]int{
				{7, 2, 5},
				{9, 9, 9},
				{1, 1, 1},
			},
			expected: []int{1, 7, 9}, // максимальные значения: 7, 9, 1 -> отсортировано: [1, 7, 9]
		},
		{
			matrix: [][]int{
				{0, -1, -5},
				{10, 10, 10},
				{2, 3, 2},
			},
			expected: []int{0, 3, 10}, // максимальные значения: 0, 10, 3 -> отсортировано: [-1, 3, 10]
		},
		{
			matrix:   [][]int{}, // пустая матрица
			expected: []int{},
		},
		{
			matrix: [][]int{
				{},
				{},
				{},
			}, // пустые строки в матрице
			expected: []int{0, 0, 0}, // результат должен быть [0, 0, 0], так как строки пустые
		},
	}

	for _, test := range tests {
		result := sliceslice(test.matrix)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("sliceslice(%v) = %v; expected %v", test.matrix, result, test.expected)
		}
	}
}
