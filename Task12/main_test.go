package main

import (
	"testing"
)

// TestConvertNumbers проверяет работу функции convertNumbers
func TestConvertNumbers(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "wki2kiaj45", // "2", "4", "5"
			expected: 11,           // 2 + 4 + 5 = 11
		},
		{
			input:    "abc123def", // "1", "2", "3"
			expected: 6,           // 1 + 2 + 3 = 6
		},
		{
			input:    "noNumbersHere", // нет чисел
			expected: 0,
		},
		{
			input:    "987xyz", // "9", "8", "7"
			expected: 24,       // 9 + 8 + 7 = 24
		},
		{
			input:    "500", // "5", "0", "0"
			expected: 5,     // 5 + 0 + 0 = 5
		},
		{
			input:    "42isAnswer", // "4", "2"
			expected: 6,            // 4 + 2 = 6
		},
		{
			input:    "", // пустая строка
			expected: 0,
		},
	}

	for _, test := range tests {
		result := convertNumbers(test.input)
		if result != test.expected {
			t.Errorf("convertNumbers(%q) = %d; expected %d", test.input, result, test.expected)
		}
	}
}
