package main

import (
	"reflect"
	"testing"
)

// TestWordsIndexes проверяет работу функции wordsIndexes
func TestWordsIndexes(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]string
	}{
		{
			input:    "hello world world hello",
			expected: map[string]string{"hello": "0 3 ", "world": "1 2 "}, // Слова "hello" и "world" с их индексами
		},
		{
			input:    "apple banana cherry",
			expected: map[string]string{"apple": "0 ", "banana": "1 ", "cherry": "2 "}, // Нет повторяющихся слов
		},
		{
			input:    "",
			expected: map[string]string{}, // Пустая строка должна вернуть пустую карту
		},
		{
			input:    "test",
			expected: map[string]string{"test": "0 "}, // Одно слово, индекс 0
		},
	}

	for _, test := range tests {
		result := wordsIndexes(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("wordsIndexes(%q) = %v; expected %v", test.input, result, test.expected)
		}
	}
}
