package main

import (
	"strings"
	"testing"
)

// TestReverseWords проверяет работу функции reverseWords
func TestReverseWords(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "Argentina manit negra",
			expected: "anitnegrA tinam argen ", // Слова перевёрнуты по отдельности
		},
		{
			input:    "hello world",
			expected: "olleh dlrow ", // "hello" -> "olleh", "world" -> "dlrow"
		},
		{
			input:    "",
			expected: "", // Пустая строка должна вернуть пустую строку
		},
		{
			input:    " test ",
			expected: "tset ", // Пробелы по краям строки игнорируются
		},
	}

	for _, test := range tests {
		result := reverseWords(test.input)
		// Удаляем лишние пробелы для корректного сравнения
		if strings.TrimSpace(result) != strings.TrimSpace(test.expected) {
			t.Errorf("reverseWords(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}
