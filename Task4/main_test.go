package main

import (
	"reflect"
	"testing"
)

// TestCommonCharacters проверяет правильность работы функции CommonCharacters
func TestCommonCharacters(t *testing.T) {
	tests := []struct {
		a, b     string
		expected []rune
	}{
		{"hello", "world", []rune{'o', 'l'}},               // общие символы 'o' и 'l'
		{"abcd", "efgh", []rune{}},                         // нет общих символов
		{"apple", "pineapple", []rune{'p', 'e', 'a', 'l'}}, // несколько общих символов, включая повторяющиеся
		{"", "abc", []rune{}},                              // одна из строк пустая
		{"abc", "", []rune{}},                              // другая строка пустая
		{"abc", "abc", []rune{'a', 'b', 'c'}},              // обе строки одинаковые
	}

	for _, test := range tests {
		result := CommonCharacters(test.a, test.b)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("CommonCharacters(%s, %s) = %v; expected %v", test.a, test.b, result, test.expected)
		}
	}
}
