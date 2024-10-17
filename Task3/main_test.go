package main

import (
	"reflect"
	"testing"
)

// TestSortString проверяет правильность работы функции SortString
func TestSortString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"listen", "eilnst"},
		{"silent", "eilnst"},
		{"google", "eggloo"},
		{"gogole", "eggloo"},
		{"", ""},
	}

	for _, test := range tests {
		result := SortString(test.input)
		if result != test.expected {
			t.Errorf("SortString(%s) = %s; expected %s", test.input, result, test.expected)
		}
	}
}

// TestGroupAnagrams проверяет правильность работы функции GroupAnagrams
func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		input    []string
		expected map[string][]string
	}{
		{
			input: []string{"listen", "silent", "enlist", "inlets", "google", "gogole"},
			expected: map[string][]string{
				"eilnst": {"listen", "silent", "enlist", "inlets"},
				"eggloo": {"google", "gogole"},
			},
		},
		{
			input: []string{"rat", "tar", "art", "star", "tars", "cheese"},
			expected: map[string][]string{
				"art":    {"rat", "tar", "art"},
				"arst":   {"star", "tars"},
				"ceehes": {"cheese"},
			},
		},
		{
			input:    []string{},
			expected: map[string][]string{},
		},
	}

	for _, test := range tests {
		result := GroupAnagrams(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("GroupAnagrams(%v) = %v; expected %v", test.input, result, test.expected)
		}
	}
}
