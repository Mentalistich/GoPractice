package main

import (
	"testing"
)

// TestFilter проверяет работу функции filter
func TestFilter(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{-3, -5, -4, 2}, 4},      // только одно положительное число: 2^2 = 4
		{[]int{1, 2, 3, 4}, 30},        // все положительные числа: 1^2 + 2^2 + 3^2 + 4^2 = 30
		{[]int{-1, -2, -3}, 0},         // все отрицательные числа удаляются
		{[]int{0, 1, -1}, 1},           // 0^2 + 1^2 = 1
		{[]int{}, 0},                   // пустой срез, результат 0
		{[]int{-10, 5, -2, 6, -7}, 61}, // оставшиеся положительные числа: 5^2 + 6^2 = 25 + 36 = 61
	}

	for _, test := range tests {
		result := filter(test.input)
		if result != test.expected {
			t.Errorf("filter(%v) = %d; expected %d", test.input, result, test.expected)
		}
	}
}
