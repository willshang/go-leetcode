package main

import "fmt"

func main() {
	// fmt.Println(strStr("hello", "ll"))
	// fmt.Println(strStr("aaaaa", "bba"))
	fmt.Println(strStr("a", "a"))
}

// leetcode 28 实现strStr()
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	next := getNext(needle)

	i := 0
	j := 0
	for i < len(haystack) && j < len(needle) {
		if j == -1 || haystack[i] == needle[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}

	if j == len(needle) {
		return i - j
	}
	return -1
}

// 求next数组
func getNext(str string) []int {
	var next = make([]int, len(str))
	next[0] = -1

	i := 0
	j := -1

	for i < len(str)-1 {
		if j == -1 || str[i] == str[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}
	return next
}
