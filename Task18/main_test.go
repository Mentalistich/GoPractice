package main

import (
	"reflect"
	"testing"
)

// TestDeleteAndReverse проверяет работу функции deleteAndReverse
func TestDeleteAndReverse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "hello world",
			expected: "drw oleh", // Удалены дубликаты: "helo wrd"; перевёрнуто: "dlrow eh"
		},
		{
			input:    "abcde",
			expected: "edcba", // Нет дубликатов, строка просто перевёрнута
		},
		{
			input:    "aabbcc",
			expected: "cba", // Удалены дубликаты: "abc"; перевёрнуто: "cba"
		},
		{
			input:    "123321",
			expected: "321", // Удалены дубликаты: "123"; перевёрнуто: "321"
		},
		{
			input:    "",
			expected: "", // Пустая строка должна вернуть пустую строку
		},
		{
			input:    "!@##@@!!",
			expected: "#@!", // Удалены дубликаты: "!@#"; перевёрнуто: "!@#"
		},
	}

	for _, test := range tests {
		result := deleteAndReverse(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("deleteAndReverse(%q) = %v; expected %v", test.input, result, test.expected)
		}
	}
}
