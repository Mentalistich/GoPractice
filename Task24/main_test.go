package main

import (
	"testing"
)

// TestFilterStrings проверяет работу функции filterStrings
func TestFilterStrings(t *testing.T) {
	tests := []struct {
		input    []string
		length   int
		expected string
	}{
		{
			input:    []string{"aew4", "weq2qe", "wqewcd", "wqewce", "gthuruhg", "eqwuheh4huehuqwe"},
			length:   6,
			expected: "wqewcd", // Строка длиной 6 после удаления цифр
		},
		{
			input:    []string{"apple123", "banana12", "cherry3"},
			length:   5,
			expected: "apple", // Строка длиной 5 после удаления цифр
		},
		{
			input:    []string{"12grape", "banana", "kiwi", "12orange34"},
			length:   6,
			expected: "banana", // Ближайшая строка к длине 6
		},
		{
			input:    []string{"short", "longer", "evenlonger"},
			length:   10,
			expected: "evenlonger", // Точная длина совпадает
		},
		{
			input:    []string{},
			length:   5,
			expected: "", // Пустой слайс должен вернуть пустую строку
		},
	}

	for _, test := range tests {
		result := filterStrings(test.input, test.length)
		if result != test.expected {
			t.Errorf("filterStrings(%v, %d) = %v; expected %v", test.input, test.length, result, test.expected)
		}
	}
}
