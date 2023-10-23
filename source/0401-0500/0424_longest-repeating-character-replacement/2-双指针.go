package main

import "fmt"

func main() {
	fmt.Println(characterReplacement("ABAB", 2))
}

func characterReplacement(s string, k int) int {
	if s == "" {
		return 0
	}
	left := 0
	count := 0
	arr := [26]int{}
	for right := 0; right < len(s); right++ {
		arr[s[right]-'A']++
		if arr[s[right]-'A'] > count {
			count = arr[s[right]-'A']
		}
		if right-left+1-count > k { // 窗口内不同字符超过k个开始收缩左边界
			arr[s[left]-'A']--
			left++
		}
	}
	return len(s) - left
}
