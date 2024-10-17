package main

import (
	"fmt"
	"strconv"
	"strings"
)

func wordsIndexes(s string) map[string]string {
	result := make(map[string]string)
	temp := strings.Fields(s)
	for i, val := range temp {
		result[val] += strconv.Itoa(i)
		result[val] += " "
	}
	return result
}

func main() {
	fmt.Println(wordsIndexes("hello world world hello"))
}
