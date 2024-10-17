package main

import "fmt"

func deleteAndReverse(s string) string {
	var temp string
	charMap := make(map[rune]bool)
	for _, value := range s {
		if !charMap[value] {
			charMap[value] = true
			temp += string(value)
		}
	}
	var result string
	for i := len(temp) - 1; i >= 0; i-- {
		result += string(temp[i])
	}
	return result
}

func main() {
	fmt.Println(deleteAndReverse("hello world"))
}
