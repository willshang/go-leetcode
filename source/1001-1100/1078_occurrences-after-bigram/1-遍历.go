package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(findOcurrences("alice is a good girl she is a good student", "a", "good"))
}

// leetcode1078_Bigram 分词
func findOcurrences(text string, first string, second string) []string {
	arr := strings.Fields(text)
	res := make([]string, 0)
	for i := 0; i < len(arr)-2; i++ {
		if arr[i] == first && arr[i+1] == second {
			res = append(res, arr[i+2])
		}
	}
	return res
}
