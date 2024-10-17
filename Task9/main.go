package main

import "fmt"

func mapMerge(map1 map[string]int, map2 map[string]int) map[string]int {
	result := make(map[string]int)
	for key, value := range map1 {
		result[key] = value
	}
	for key, val := range map2 {
		if _, ok := result[key]; ok {
			if result[key]+val > 10 {
				result[key] += val
			} else {
				delete(result, key)
			}
		} else {
			result[key] = val
		}

	}

	return result
}

func main() {
	map1 := map[string]int{
		"apple":  5,
		"banana": 4,
	}

	map2 := map[string]int{
		"banana": 4,
		"cherry": 3,
	}
	fmt.Println(mapMerge(map1, map2))
}
