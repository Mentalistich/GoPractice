package main

import (
	"fmt"
	"sort"
)

func superEffectiveFunc(s []string) string {
	var result string
	for i, val := range s {
		for j := range s {
			if val == s[j] && i != j {
				s[j] = ""
			}
		}
	}
	for i := 0; i < len(s); i++ {
		if s[i] == "" {
			s = append(s[:i], s[i+1:]...)
			i--
		}
	}
	sort.Slice(s, func(i, j int) bool { return len(s[i]) < len(s[j]) })
	for i := range s {
		result += s[i]
		if i != len(s)-1 {
			result += ","
		}
	}
	return result
}
func main() {
	test := []string{"sqeqw", "wqes", "sqe", "sqe"}
	fmt.Println(superEffectiveFunc(test))
}
