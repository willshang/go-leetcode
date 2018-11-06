package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}
func reverseWords(s string) string {
	strs := strings.Split(s," ")
	for i, s := range strs{
		strs[i] = reverse(s)
	}
	return strings.Join(strs," ")
}

func reverse(s string) string  {
	bytes := []byte(s)
	i, j := 0, len(bytes)-1
	for i < j{
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
	return string(bytes)
}