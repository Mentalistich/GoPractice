package main

import (
	"fmt"
	"strconv"
)

func convertNumbers(s string) int {
	var result int
	for _, val := range s {
		if value, err := strconv.Atoi(string(val)); err != nil {
			continue
		} else {
			result += value
		}
	}
	return result
}

func main() {
	s := "wki2kiaj45"
	fmt.Println(convertNumbers(s))
}
