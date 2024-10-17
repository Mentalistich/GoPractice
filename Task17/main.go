package main

import "fmt"

func average(slice []int) []int {
	temp := make([]int, 0)
	for i := 0; i <= len(slice)-1; i += 2 {
		if i == len(slice)-1 {
			temp = append(temp, slice[i])
		} else {
			temp = append(temp, slice[i]+slice[i+1])
		}
	}
	var aver int
	for _, val := range temp {
		aver += val
	}
	if len(temp) != 0 {
		aver = aver / len(temp)
		fmt.Println(aver)
	} else {

		return temp
	}
	result := make([]int, 0)
	for _, val := range temp {
		if val > aver {
			result = append(result, val)
		}
	}
	return result
}

func main() {
	fmt.Println(average([]int{5, 5, 5, 5, 5}))
}
