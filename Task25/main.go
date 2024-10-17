package main

import (
	"fmt"
	"sort"
)

func newSlice(slice []int) []int {
	result := make([]int, 0)
	for i, value := range slice {
		if i%2 == 0 {
			result = append(result, value*2)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return (result[i]) > (result[j])
	})
	return result
}

func main() {
	fmt.Println(newSlice([]int{1, 2, 3, 4, 5}))
}
