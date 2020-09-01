package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minWindow("ADOBECODEBANCDABC", "ABC"))
}

// leetcode76_最小覆盖子串
func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	window := make(map[byte]int)
	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	left, right := -1, -1
	minLength := math.MaxInt32
	for l, r := 0, 0; r < len(s); r++ {
		if r < len(s) && need[s[r]] > 0 {
			window[s[r]]++
		}
		// 找到，然后left往右移
		for check(need, window) == true && l <= r {
			if r-l+1 < minLength {
				minLength = r - l + 1
				left, right = l, r+1
			}
			if _, ok := need[s[l]]; ok {
				window[s[l]]--
			}
			l++
		}
	}
	if left == -1 {
		return ""
	}
	return s[left:right]
}

func check(need, window map[byte]int) bool {
	for k, v := range need {
		if window[k] < v {
			return false
		}
	}
	return true
}
