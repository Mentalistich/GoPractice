package main

import (
	"reflect"
	"testing"
)

// TestNewSlice проверяет работу функции newSlice
func TestNewSlice(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{10, 6, 2}, // Умножаем элементы с индексами 0, 2, 4: [2, 6, 10], затем сортируем по убыванию
		},
		{
			input:    []int{2, 3, 4, 5, 6, 7},
			expected: []int{12, 8, 4}, // Умножаем элементы с индексами 0, 2, 4: [4, 8, 12], сортируем по убыванию
		},
		{
			input:    []int{1},
			expected: []int{2}, // Один элемент, умножаем его на 2
		},
		{
			input:    []int{},
			expected: []int{}, // Пустой слайс, должен вернуть пустой результат
		},
		{
			input:    []int{10, 20, 30, 40, 50},
			expected: []int{100, 60, 20}, // Умножаем элементы с индексами 0, 2, 4: [20, 60, 100], сортируем по убыванию
		},
	}

	for _, test := range tests {
		result := newSlice(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("newSlice(%v) = %v; expected %v", test.input, result, test.expected)
		}
	}
}
