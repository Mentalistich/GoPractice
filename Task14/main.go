package main

import (
	"fmt"
	"sort"
	"strings"
)

// Группировка и сортировка слов по длине:
// Напишите функцию, которая принимает строку, разбивает её на слова, группирует слова по их длине в map (ключ — длина слова, значение — список слов),
// а затем возвращает слайс этих групп, отсортированных по длине.
func stringsLen(s string) [][]string {
	str := strings.Split(s, " ")
	res := make(map[int][]string)
	result := make([][]string, 0)
	for _, val := range str {
		res[len(val)] = append(res[len(val)], val)
	}
	for _, val := range res {
		result = append(result, val)
	}
	sort.Slice(result, func(i, j int) bool {
		return len(result[i]) < len(result[j])
	})
	return result

}

func main() {
	s := "apple apple apple apple bat cat dog elephant frog"
	fmt.Println(stringsLen(s))
}
