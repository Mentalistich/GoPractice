package main

import (
	"reflect"
	"testing"
)

// TestDeleteAllPunctuation проверяет работу функции deleteAllPunctuation
func TestDeleteAllPunctuation(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]int
	}{
		{
			input: "WEWQ!wewq,wqueqwuje,.qwuehqwueh",
			expected: map[string]int{
				"wewq":       2,
				"wqueqwuje":  1,
				"qwuehqwueh": 1,
			},
			// Пунктуация удалена, строка переведена в нижний регистр, слова подсчитаны
		},
		{
			input: "Hello, world! This is a test.",
			expected: map[string]int{
				"hello": 1,
				"world": 1,
				"this":  1,
				"is":    1,
				"a":     1,
				"test":  1,
			},
			// Слова "Hello" и "world" после удаления пунктуации и перевода в нижний регистр
		},
		{
			input: "NoPunctuationHere",
			expected: map[string]int{
				"nopunctuationhere": 1,
			},
			// Строка без пунктуации, просто приводится к нижнему регистру
		},
		{
			input: "word!!!",
			expected: map[string]int{
				"word": 1,
			},
			// Одно слово с пунктуацией
		},
		{
			input:    "", // Пустая строка
			expected: map[string]int{},
		},
	}

	for _, test := range tests {
		result := deleteAllPunctuation(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("deleteAllPunctuation(%q) = %v; expected %v", test.input, result, test.expected)
		}
	}
}
