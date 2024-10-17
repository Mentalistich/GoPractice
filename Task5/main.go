package main

import (
	"fmt"
	"strconv"
)

func CompressString(s string) string {
	var str string
	var counter int = 1
	for i, val := range s {
		if len(s) > i+1 && val == rune(s[i+1]) {
			counter++
		} else {
			if len(s)-1 == i+1 && val != rune(s[i]) {
			} else {
				str += string(val)
				if counter > 1 {
					str += strconv.Itoa(counter)
				}
				counter = 1
			}
		}
	}

	return str
}
func main() {
	fmt.Println(CompressString("aaabccccsa"))
}
