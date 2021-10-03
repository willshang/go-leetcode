package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minWindow("ADOBECODEBANCDABC", "ABC"))
}

// 剑指OfferII017.含有所有字符的最短字符串
func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	arr := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		arr[t[i]]++
	}
	l, count := 0, 0
	res := ""
	minLength := math.MaxInt32
	for r := 0; r < len(s); r++ {
		arr[s[r]]--
		if arr[s[r]] >= 0 {
			count++
		}
		// left往右边移动
		for count == len(t) {
			if minLength > r-l+1 {
				minLength = r - l + 1
				res = s[l : r+1]
			}
			arr[s[l]]++
			if arr[s[l]] > 0 {
				count--
			}
			l++
		}
	}
	return res
}
