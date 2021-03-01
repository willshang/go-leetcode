package main

import "fmt"

func main() {
	fmt.Println(characterReplacement("ABAB", 2))
}

// leetcode424_替换后的最长重复字符
func characterReplacement(s string, k int) int {
	if s == "" {
		return 0
	}
	res := 0
	left := 0
	count := 0
	arr := [26]int{}
	for right := 0; right < len(s); right++ {
		arr[s[right]-'A']++
		if arr[s[right]-'A'] > count {
			count = arr[s[right]-'A']
		}
		for right-left+1-count > k {
			arr[s[left]-'A']--
			left++
		}
		if right-left+1 > res {
			res = right - left + 1
		}
	}
	return res
}
