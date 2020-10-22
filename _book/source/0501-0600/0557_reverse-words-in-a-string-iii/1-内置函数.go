package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}

func reverseWords(s string) string {
	strS := strings.Split(s, " ")
	for i, s := range strS {
		strS[i] = reverse(s)
	}
	return strings.Join(strS, " ")
}

func reverse(s string) string {
	arr := []byte(s)
	i, j := 0, len(arr)-1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	return string(arr)
}
