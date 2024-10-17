package main

import (
	"reflect"
	"testing"
)

// TestSortMap проверяет работу функции sortMap
func TestSortMap(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{
			input:    "hello world go test case",
			expected: []string{"go", "test", "case", "hello", "world"}, // Слова отсортированы по длине
		},
		{
			input:    "apple banana cherry apple banana",
			expected: []string{"apple", "banana", "cherry"}, // Повторяющиеся слова удалены
		},
		{
			input:    "",
			expected: []string{}, // Пустая строка должна вернуть пустой срез
		},
		{
			input:    "one",
			expected: []string{"one"}, // Строка с одним словом
		},
	}

	for _, test := range tests {
		result := sortMap(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("sortMap(%q) = %v; expected %v", test.input, result, test.expected)
		}
	}
}
