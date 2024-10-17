package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func filterStrings(s []string, lenght int) string {
	var tempStr string
	var resultStr string
	for _, val := range s {
		for _, letter := range val {
			if !unicode.IsDigit(letter) {
				tempStr += string(letter)
			}
		}
		tempStr += " "
	}
	result := strings.Fields(tempStr)
	sort.Slice(result, func(i, j int) bool {
		return len(result[i]) < len(result[j])
	})
	minimum := 10000
	var indexesMin int
	for i, value := range result {
		if len(value) > lenght {
			if minimum > len(value)-lenght {
				minimum = len(value) - lenght
				indexesMin = i
			}
		}
		if len(value) == lenght && len(resultStr) == 0 {
			resultStr = value
		}
	}
	if len(resultStr) == 0 && len(result) != 0 {
		return result[indexesMin]
	} else {
		return resultStr
	}

}
func main() {
	fmt.Println(filterStrings([]string{"aew4", "weq2qe", "wqewcd", "wqewce", "gthuruhg", "eqwuheh4huehuqwe"}, 6))

}
