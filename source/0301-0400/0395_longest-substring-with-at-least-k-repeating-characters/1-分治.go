package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(longestSubstring("ababacb", 3))
}

func longestSubstring(s string, k int) int {
	res := 0
	m := make(map[byte]int) // 统计每个字符的个数
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	var split byte // 某个字符出现的次数：0<x<k
	for key, value := range m {
		if value < k {
			split = key
			break
		}
	}
	if split == 0 { // 字符出现次数都>=k
		return len(s)
	}
	arr := strings.Split(s, string(split)) // 以该字符切割的子串
	for i := 0; i < len(arr); i++ {
		temp := longestSubstring(arr[i], k)
		if temp > res {
			res = temp
		}
	}
	return res
}
