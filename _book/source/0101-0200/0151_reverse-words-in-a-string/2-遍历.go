package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("the sky is blue"))
}

func reverseWords(s string) string {
	arr := make([]string, 0)
	i, j := 0, 0
	for i < len(s) && j <= len(s) {
		for i = j; i < len(s) && s[i] == ' '; i++ {
		}
		for j = i; j < len(s) && s[j] != ' '; j++ {
		}
		if i < j {
			arr = append(arr, s[i:j])
		}
	}
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	return strings.Join(arr, " ")
}
