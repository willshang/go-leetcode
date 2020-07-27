package main

import "fmt"

func main() {
	fmt.Println(reverseLeftWords("abcdefg", 2))
}

// 剑指offer_面试题58-II_左旋转字符串
func reverseLeftWords(s string, n int) string {
	n = n % len(s)
	arr := []byte(s)
	reverse(arr, 0, n-1)
	reverse(arr, n, len(arr)-1)
	reverse(arr, 0, len(arr)-1)
	return string(arr)
}

func reverse(arr []byte, start, end int) []byte {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
	return arr
}
