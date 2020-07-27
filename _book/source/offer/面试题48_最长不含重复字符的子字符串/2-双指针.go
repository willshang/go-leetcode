package main

import "fmt"

func main() {
	//fmt.Println(lengthOfLongestSubstring("abcabcabcbb"))
	fmt.Println(lengthOfLongestSubstring("aabaab!bb"))
}

func lengthOfLongestSubstring(s string) int {
	if len(s) < 1 {
		return 0
	}
	left := 0
	right := 0
	res := 0
	for left <= right {
		m := make(map[byte]int)
		for i := left; i < right; i++ {
			m[s[i]]++
		}
		for right < len(s) {
			if value, ok := m[s[right]]; ok && value > 0 {
				if right-left > res {
					res = right - left
				}
				left++
				break
			} else {
				m[s[right]]++
				right++
			}
		}
		if right-left > res {
			res = right - left
		}
		if right >= len(s)-1 {
			break
		}
	}
	return res
}
