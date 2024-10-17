package main

import (
	"reflect"
	"testing"
)

// TestMapMerge проверяет работу функции mapMerge
func TestMapMerge(t *testing.T) {
	tests := []struct {
		map1     map[string]int
		map2     map[string]int
		expected map[string]int
	}{
		{
			map1: map[string]int{
				"apple":  5,
				"banana": 4,
			},
			map2: map[string]int{
				"banana": 4,
				"cherry": 3,
			},
			expected: map[string]int{
				"apple":  5,
				"cherry": 3,
			}, // banana is removed (4 + 4 = 8, less than 10)
		},
		{
			map1: map[string]int{
				"apple":  6,
				"banana": 6,
			},
			map2: map[string]int{
				"banana": 5,
				"pear":   2,
			},
			expected: map[string]int{
				"apple":  6,
				"banana": 11,
				"pear":   2,
			}, // banana is kept (6 + 5 = 11, greater than 10)
		},
		{
			map1: map[string]int{
				"mango":  7,
				"orange": 2,
			},
			map2: map[string]int{
				"mango": 4,
				"kiwi":  5,
			},
			expected: map[string]int{
				"kiwi":   5,
				"mango":  11,
				"orange": 2,
			}, // mango is kept (7 + 4 = 11, greater than 10)
		},
		{
			map1:     map[string]int{}, // empty map1
			map2:     map[string]int{"grape": 5},
			expected: map[string]int{"grape": 5}, // only map2 keys
		},
		{
			map1: map[string]int{"grape": 6},
			map2: map[string]int{}, // empty map2
			expected: map[string]int{
				"grape": 6, // only map1 keys
			},
		},
		{
			map1: map[string]int{
				"apple":  5,
				"banana": 6,
			},
			map2: map[string]int{
				"banana": 4,
				"apple":  7,
			},
			expected: map[string]int{
				"apple": 12, // apple is kept (5 + 7 = 12, greater than 10)
			}, // banana is removed (6 + 4 = 10, equal to 10)
		},
	}

	for _, test := range tests {
		result := mapMerge(test.map1, test.map2)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("mapMerge(%v, %v) = %v; expected %v", test.map1, test.map2, result, test.expected)
		}
	}
}
