package main

import (
	"reflect"
	"testing"
)

// TestReturnMap проверяет работу функции returnMap
func TestReturnMap(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]int
	}{
		{
			input:    "qqq www sss qqq sss",
			expected: map[string]int{"qqq": 2, "sss": 2}, // Повторяются "qqq" и "sss"
		},
		{
			input:    "apple banana cherry",
			expected: map[string]int{}, // Нет повторяющихся слов
		},
		{
			input:    "",
			expected: map[string]int{}, // Пустая строка должна вернуть пустую карту
		},
		{
			input:    "test test test test",
			expected: map[string]int{"test": 4}, // Одно повторяющееся слово "test"
		},
	}

	for _, test := range tests {
		result := returnMap(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("returnMap(%q) = %v; expected %v", test.input, result, test.expected)
		}
	}
}
