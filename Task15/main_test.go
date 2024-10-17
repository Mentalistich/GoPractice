package main

import (
	"reflect"
	"testing"
)

// TestUpRegisterStrings проверяет работу функции upRegisterStrings
func TestUpRegisterStrings(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{
			input:    []string{"qqq", "wsdadd", "qweqweokqwi", "weiqjeiwqjieij", "wqj", "wqejiwiejiq", "ewqj"},
			expected: []string{"WQEJIWIEJIQ", "WEIQJEIWQJIEIJ", "QWEQWEOKQWI", "WSDADD"},
			// Слова длиной больше 5 символов: "wsdadd", "qweqweokqwi", "weiqjeiwqjieij", "wqejiwiejiq" преобразованы в верхний регистр и возвращены в обратном порядке.
		},
		{
			input:    []string{"short", "tiny", "ok", "small"},
			expected: []string{}, // Все строки имеют длину меньше 5 символов.
		},
		{
			input:    []string{"longword", "anotherlongword"},
			expected: []string{"ANOTHERLONGWORD", "LONGWORD"}, // Обе строки длиннее 5 символов, результат в обратном порядке.
		},
		{
			input:    []string{}, // Пустой срез
			expected: []string{}, // Ожидается пустой срез.
		},
		{
			input:    []string{"equal", "equal", "equal"}, // Все строки одинаковой длины (менее 5 символов).
			expected: []string{},                          // Все строки меньше 5 символов.
		},
	}

	for _, test := range tests {
		result := upRegisterStrings(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("upRegisterStrings(%v) = %v; expected %v", test.input, result, test.expected)
		}
	}
}
