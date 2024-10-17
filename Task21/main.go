package main

import "fmt"

func multipleSlice(slice []int) []int {
	result := make([]int, 0)
	var average int
	for _, val := range slice {
		average += val
	}
	for i, value := range slice {
		if value > 0 {
			if value*i > average {
				result = append(result, value)
			}
		}
	}
	return result
}

func main() {
	fmt.Println(multipleSlice([]int{1, 2, 3, 4}))
}
