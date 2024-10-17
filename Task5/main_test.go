package main

import (
	"testing"
)

// TestCompressString проверяет работу функции CompressString
func TestCompressString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"aaabccccsa", "a3bc4sa"}, // несколько повторяющихся символов
		{"abcd", "abcd"},          // нет повторяющихся символов
		{"aabbcc", "a2b2c2"},      // все символы повторяются дважды
		{"", ""},                  // пустая строка
		{"a", "a"},                // одиночный символ
		{"aaa", "a3"},             // все символы одинаковые
		{"aabbbaa", "a2b3a2"},     // повторяющиеся блоки символов
	}

	for _, test := range tests {
		result := CompressString(test.input)
		if result != test.expected {
			t.Errorf("CompressString(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}
