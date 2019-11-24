package main

import "fmt"

func main() {
	str := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(str))
}
func lengthOfLongestSubstring(s string) int {
	location := [256]int{}
	for i := range location {
		location[i] = -1
	}

	maxLen, left := 0, 0

	for i := 0; i < len(s); i++ {
		fmt.Println(left,maxLen)
		if location[s[i]] >= left {
			//fmt.Println("重复了:\t",s[i:],string(s[i]),location[s[i]],left)
			left = location[s[i]] + 1
		} else if i+1-left > maxLen {
			maxLen = i + 1 - left
		}

		location[s[i]] = i
	}

	//for k, v := range location {
	//	if v != -1 {
	//		fmt.Println(k, v)
	//	}
	//}
	return maxLen
}
