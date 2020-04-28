package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}

// leetcode557_反转字符串中的单词 III
func reverseWords(s string) string {
	arr := []byte(s)
	j := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == ' ' {
			reverse(arr, j, i-1)
			j = i + 1
		}
	}
	reverse(arr, j, len(arr)-1)
	return string(arr)
}

func reverse(arr []byte, i, j int) []byte {
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	return arr
}
