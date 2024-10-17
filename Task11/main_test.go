package main

import (
	"reflect"
	"testing"
)

// TestSuperEffectiveFunc проверяет работу функции superEffectiveFunc
func TestSuperEffectiveFunc(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{
			input:    []string{"sqeqw", "wqes", "sqe", "sqe"},
			expected: "sqe,wqes,sqeqw", // "sqe" дублируется, удаляется, затем сортировка по длине
		},
		{
			input:    []string{"apple", "banana", "cherry", "banana", "apple"},
			expected: "apple,banana,cherry", // дубликаты "apple" и "banana" удаляются
		},
		{
			input:    []string{"dog", "cat", "elephant", "tiger", "dog", "cat"},
			expected: "dog,cat,tiger,elephant", // дубликаты "dog" и "cat" удаляются
		},
		{
			input:    []string{"a", "abc", "ab", "abcd", "abc", "ab"},
			expected: "a,ab,abc,abcd", // дубликаты "abc" и "ab" удаляются, сортировка по длине
		},
		{
			input:    []string{}, // пустой срез
			expected: "",
		},
		{
			input:    []string{"single"},
			expected: "single", // один элемент, нет дубликатов
		},
	}

	for _, test := range tests {
		result := superEffectiveFunc(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("superEffectiveFunc(%v) = %v; expected %v", test.input, result, test.expected)
		}
	}
}
