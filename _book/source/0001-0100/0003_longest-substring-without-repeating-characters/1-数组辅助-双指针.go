package main

import "fmt"

func main() {
	//fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring(" "))
}

func lengthOfLongestSubstring(s string) int {
	arr := [256]int{}
	for i := range arr {
		arr[i] = -1
	}
	max, j := 0, 0
	for i := 0; i < len(s); i++ {
		if arr[s[i]] >= j {
			j = arr[s[i]] + 1
		} else if i+1-j > max {
			max = i + 1 - j
		}
		arr[s[i]] = i
	}
	return max
}
