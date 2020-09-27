package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(arrangeWords("Leetcode is cool"))
	fmt.Println(arrangeWords("gbo ops"))
}

// leetcode1451_重新排列句子中的单词
func arrangeWords(text string) string {
	text = strings.ToLower(text)
	arr := strings.Fields(text)
	sort.SliceStable(arr, func(i, j int) bool {
		return len(arr[i]) < len(arr[j])
	})
	arr[0] = strings.Title(arr[0])
	return strings.Join(arr, " ")
}
