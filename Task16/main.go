package main

import (
	"fmt"
	"strings"
	"unicode"
)

func deleteAllPunctuation(s string) map[string]int {
	var str string
	s = strings.ToLower(s)
	for _, let := range s {
		if unicode.IsPunct(let) {
			str += " "
		} else {
			str += string(let)
		}
	}
	letters := strings.Fields(str)
	result := make(map[string]int)
	for _, val := range letters {
		result[val]++
	}
	return result

}
func main() {
	fmt.Println(deleteAllPunctuation("WEWQ!wewq,wqueqwuje,.qwuehqwueh"))
}
