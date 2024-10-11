package main

import (
	"fmt"
	"sort"
)

func sliceslice(matrix [][]int) []int {
	result := make([]int, len(matrix))
	for i, val := range matrix {
		if len(val) > 0 {
			max := val[0]
			for _, value := range val {
				if value > max {
					max = value
				}
			}
			result[i] = max
		} else {
			result[i] = 0
		}
	}
	sort.Ints(result)
	return result
}
func main() {
	slices := [][]int{
		{3, 3, 1},
		{4, 0, 4},
		{3, 1, 1},
	}

	fmt.Println(sliceslice(slices))
}
