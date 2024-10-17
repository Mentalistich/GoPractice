package main

import (
	"reflect"
	"testing"
)

// TestMultipleSlice проверяет работу функции multipleSlice
func TestMultipleSlice(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{1, 2, 3, 4},
			expected: []int{4}, // Элементы, где value * index > average
		},
		{
			input:    []int{-1, -2, -3, -4},
			expected: []int{}, // Все отрицательные числа, ни одно из них не удовлетворяет условию
		},
		{
			input:    []int{},
			expected: []int{}, // Пустой срез должен вернуть пустой результат
		},
		{
			input:    []int{100},
			expected: []int{}, // Единственный элемент не может удовлетворить условию (0 * 100 = 0)
		},
	}

	for _, test := range tests {
		result := multipleSlice(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("multipleSlice(%v) = %v; expected %v", test.input, result, test.expected)
		}
	}
}
