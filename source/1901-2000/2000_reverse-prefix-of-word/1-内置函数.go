package main

import "strings"

func main() {

}

// leetcode2000_反转单词前缀
func reversePrefix(word string, ch byte) string {
	index := strings.Index(word, string(ch))
	arr := []byte(word)
	reverse(arr[:index+1])
	return string(arr)
}

func reverse(arr []byte) []byte {
	start := 0
	end := len(arr) - 1
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
	return arr
}
