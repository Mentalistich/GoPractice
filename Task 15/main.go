package main

import (
	"fmt"
	"strings"
)

func upRegisterStrings(s []string) []string {
	temp := make([]string, 0)
	for _, val := range s {
		if len(val) > 5 {
			temp = append(temp, strings.ToUpper(val))
		}
	}
	result := make([]string, len(temp))
	for i := 0; i < len(temp); i++ {
		for i := 0; i < len(temp); i++ {
			// Копируем элементы в обратном порядке
			result[i] = temp[len(temp)-1-i]
		}
	}
	return result
}

func main() {
	test := []string{"qqq", "wsdadd", "qweqweokqwi", "weiqjeiwqjieij", "wqj", "wqejiwiejiq", "ewqj"}
	fmt.Println(upRegisterStrings(test))
}
