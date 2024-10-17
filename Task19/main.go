package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortMap(s string) []string {
	temp := strings.Fields(s)
	myMap := make(map[string]int)
	result := make([]string, 0)
	sort.Slice(temp, func(i, j int) bool {
		return len(temp[i]) < len(temp[j])
	})

	for _, val := range temp {
		myMap[val] = len(val)
	}
	for key := range myMap {
		result = append(result, key)
	}
	sort.Slice(result, func(i, j int) bool {
		return len(result[i]) < len(result[j])
	})
	return result
}

func main() {
	fmt.Println(sortMap("wqek uwqhequwhe qwuheqwhuehu wqek"))
}
