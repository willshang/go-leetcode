package main

import "fmt"

func main() {
	fmt.Println(reverseStr("abcdefg", 2))
}

// leetcode541_反转字符串II
func reverseStr(s string, k int) string {
	arr := []byte(s)
	for i := 0; i < len(s); i = i + 2*k {
		j := min(i+k, len(s))
		reverse(arr[i:j])
	}
	return string(arr)
}

func reverse(arr []byte) {
	i, j := 0, len(arr)-1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
