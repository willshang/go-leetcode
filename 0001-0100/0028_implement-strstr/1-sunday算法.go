package main

import "fmt"

func main() {
	// fmt.Println(strStr("hello", "ll"))
	// fmt.Println(strStr("aaaaa", "bba"))
	fmt.Println(strStr("a", "a"))
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	// 01.计算模式串needle的偏移量
	m := make(map[int32]int)
	for k, v := range needle {
		m[v] = len(needle) - k
	}

	index := 0
	for index+len(needle) <= len(haystack) {
		// 匹配字符串
		str := haystack[index : index+len(needle)]
		if str == needle {
			return index
		} else {
			if index+len(needle) >= len(haystack) {
				return -1
			}
			// 后一位字符串
			next := haystack[index+len(needle)]
			if nextStep, ok := m[int32(next)]; ok {
				index = index + nextStep
			} else {
				index = index + len(needle) + 1
			}
		}
	}
	if index+len(needle) >= len(haystack) {
		return -1
	} else {
		return index
	}
}
