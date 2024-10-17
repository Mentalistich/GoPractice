package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	for i, value := range r {
		r[i] = unicode.ToLower(value)
	}
	sort.Sort(sortRunes(r))
	return string(r)
}

func GroupAnagrams(words []string) map[string][]string {
	group := make(map[string][]string)
	for i, val := range words {
		// TODO разобраться как это устроено
		//sort.Slice(val, func(i, j int) bool {
		//	return val[i] < val[j]
		//})
		val = SortString(strings.ToLower(val))
		group[val] = append(group[val], words[i])
	}
	fmt.Println(group)
	return group
}
func main() {
	var words []string = []string{"listen", "silent", "enlist", "inlets", "google", "gogole"}
	GroupAnagrams(words)
}
