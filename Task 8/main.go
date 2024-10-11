package main

import "fmt"

func filter(slice []int) int {
	var result int
	for i := 0; i < len(slice); i++ {
		if slice[i] < 0 {
			slice = append(slice[:i], slice[i+1:]...)
			i--
		} else {
			slice[i] *= slice[i]
		}
	}
	for _, val := range slice {
		result += val
	}
	return result
}

func main() {
	x := []int{-3, -5, -4, 2}
	fmt.Println(filter(x))
}
