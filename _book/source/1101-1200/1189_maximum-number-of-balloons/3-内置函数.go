package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(maxNumberOfBalloons("loonbalxballpoon"))
}

func maxNumberOfBalloons(text string) int {
	arr := make([]int, 0)
	str := "ablon"
	for i := 0; i < len(str); i++ {
		if str[i] == 'l' || str[i] == 'o' {
			arr = append(arr, strings.Count(text, string(str[i]))/2)
		} else {
			arr = append(arr, strings.Count(text, string(str[i])))
		}
	}
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}
