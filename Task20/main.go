package main

import (
	"fmt"
	"strings"
)

func returnMap(s string) map[string]int {
	temp := strings.Fields(s)
	tempMap := make(map[string]int)
	result := make(map[string]int)
	for _, value := range temp {
		tempMap[value]++
	}
	for key, value := range tempMap {
		if tempMap[key] > 1 {
			result[key] = value
		}
	}
	return result
}

func main() {
	fmt.Println(returnMap("qqq www sss qqq sss "))
}
