package main

import (
	"reflect"
	"testing"
)

// TestAverage проверяет работу функции average
func TestAverage(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{1, 2, 1, 2, 100, 100},
			expected: []int{200}, // Суммы: [3, 3, 200]; Среднее: 68; Число больше среднего: 200
		},
		{
			input:    []int{10, 20, 30, 40, 50, 60},
			expected: []int{110}, // Суммы: [30, 70, 110]; Среднее: 70; Числа больше среднего: 70, 110
		},
		{
			input:    []int{5, 5, 5, 5, 5}, // Нечётное количество элементов
			expected: []int{10, 10},        // Суммы: [10, 10, 5]; Среднее: 8; Число больше среднего: 10
		},
		{
			input:    []int{100, 100, 100, 100}, // Все элементы одинаковы
			expected: []int{},                   // Суммы: [200, 200]; Среднее: 200; Нет элементов больше среднего
		},
		{
			input:    []int{}, // Пустой список
			expected: []int{}, // Пустой результат
		},
	}

	for _, test := range tests {
		result := average(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("average(%v) = %v; expected %v", test.input, result, test.expected)
		}
	}
}
