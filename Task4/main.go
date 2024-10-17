package main

import (
	"fmt"
	"slices"
	"strings"
)

func CommonCharacters(a, b string) []rune {
	var eqRune []rune
	for _, val := range a {
		for _, value := range b {
			if strings.ContainsRune(a, val) && strings.ContainsRune(a, value) && !slices.Contains(eqRune, value) {
				eqRune = append(eqRune, value)
			}
		}
	}
	return eqRune
}
func main() {
	a := CommonCharacters("apple", "pineapple")
	fmt.Println(string(a))
}
