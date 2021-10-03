package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	//fmt.Println(lengthOfLongestSubstring("tmmzuxt"))
	//fmt.Println(lengthOfLongestSubstring(" "))
}

func lengthOfLongestSubstring(s string) int {
	m := make(map[uint8]int)
	max, j := 0, 0
	for i := 0; i < len(s); i++ {
		if v, ok := m[s[i]]; ok && v >= j {
			j = v + 1
		} else if i+1-j > max {
			max = i + 1 - j
		}
		m[s[i]] = i
	}
	return max
}
