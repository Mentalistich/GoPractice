package main

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
func stringInfo(s string) map[string]int {
	result := make(map[string]int)
	str := strings.Fields(s)
	var indexes []int
	for i := range str {
		if len(str[i]) < 3 {
			indexes = append(indexes, i)
		}
	}
	sort.Slice(indexes, func(i, j int) bool { return indexes[i] > indexes[j] })
	for i := len(str) - 2; i >= 0; i-- {
		if slices.Contains(indexes, i) {
			str = remove(str, i)
		}
	}
	for i := range str {
		if len(str[i]) < 3 {
			str = remove(str, i)
		}
	}
	//TODO почему не работает этот вариант?
	//if len(str[len(str)-2]) < 3 {
	//	str = remove(str, len(str)-2)
	//}
	for _, val := range str {
		if _, ok := result[val]; ok {
			result[val]++
		} else {
			result[val] = 1
		}

	}
	return result
}

func main() {
	fmt.Println(stringInfo("hello hello ewq xasqwe e sdasd ew world w w ws"))
}
