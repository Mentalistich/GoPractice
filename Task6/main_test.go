package main

import (
	"reflect"
	"testing"
)

// TestStringInfo проверяет правильность работы функции stringInfo
func TestStringInfo(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]int
	}{
		{
			input: "hello hello ewq xasqwe e sdasd ew world w w ws",
			expected: map[string]int{
				"hello":  2,
				"ewq":    1,
				"xasqwe": 1,
				"sdasd":  1,
				"world":  1,
			},
		},
		{
			input: "a bb ccc ddd e fff ggg h",
			expected: map[string]int{
				"ccc": 1,
				"ddd": 1,
				"fff": 1,
				"ggg": 1,
			},
		},
		{
			input: "word1 word2 word1 word3 word3",
			expected: map[string]int{
				"word1": 2,
				"word2": 1,
				"word3": 2,
			},
		},
		{
			input:    "a b c", // все слова короче 3 символов
			expected: map[string]int{},
		},
		{
			input:    "", // пустая строка
			expected: map[string]int{},
		},
	}

	for _, test := range tests {
		result := stringInfo(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("stringInfo(%q) = %v; expected %v", test.input, result, test.expected)
		}
	}
}
