package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestSubstring("ababacb", 3))
}

// leetcode395_至少有K个重复字符的最长子串
func longestSubstring(s string, k int) int {
	res := 0
	// 多次滑动窗口
	for i := 1; i < 26; i++ { // 枚举字符种类的个数
		m := make(map[byte]int)
		count := 0
		left := 0
		lessK := 0 // 字符出现次数<k的个数
		for right := 0; right < len(s); right++ {
			if m[s[right]] == 0 {
				count++
				lessK++
			}
			m[s[right]]++
			if m[s[right]] == k {
				lessK--
			}
			for count > i { // 字母出现次数大于枚举的次数
				if m[s[left]] == k {
					lessK++
				}
				m[s[left]]-- // 移动左窗口
				if m[s[left]] == 0 {
					count--
					lessK--
				}
				left++
			}
			if lessK == 0 && right-left+1 > res { // 更新结果
				res = right - left + 1
			}
		}
	}
	return res
}
