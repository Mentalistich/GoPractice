package main

import (
	"reflect"
	"testing"
)

// TestStringsLen проверяет работу функции stringsLen
func TestStringsLen(t *testing.T) {
	tests := []struct {
		input    string
		expected [][]string
	}{

		{
			input:    "hello world go",
			expected: [][]string{{"go"}, {"hello", "world"}},
			// Слова "hello" и "world" сгруппированы по длине 5, "go" по длине 2
		},
		{
			input:    "a bb ccc dddd eeeee",
			expected: [][]string{{"a"}, {"bb"}, {"ccc"}, {"dddd"}, {"eeeee"}},
			// Каждое слово имеет свою уникальную длину
		},
		{
			input:    "same same same same",
			expected: [][]string{{"same", "same", "same", "same"}},
			// Все слова одинаковы, длина у них одна
		},
		{
			input:    "", // Пустая строка
			expected: [][]string{},
		},
		{
			input:    "word", // Строка с одним словом
			expected: [][]string{{"word"}},
		},
	}

	for _, test := range tests {
		result := stringsLen(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("stringsLen(%q) = %v; expected %v", test.input, result, test.expected)
		}
	}
}
