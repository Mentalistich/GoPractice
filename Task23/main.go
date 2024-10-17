package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	temp := strings.Fields(s)
	var str string
	var reverseStr string

	for _, val := range temp {
		str = val
		for j := len(str) - 1; j >= 0; j-- {
			reverseStr += string(str[j])
		}
		reverseStr += " "
	}
	return reverseStr
}

func main() {
	fmt.Println(reverseWords("Argentina manit negra"))
}
