package main

import "fmt"

func main() {
	fmt.Println(maxFreq("aaaa", 1, 3, 3))
}

func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	res := 0
	m := make(map[string]int)
	window := make(map[byte]int)
	count := 0
	left := 0
	for i := 0; i < len(s); i++ {
		window[s[i]]++
		if window[s[i]] == 1 {
			count++
		}
		length := i - left + 1
		if length < minSize {
			continue
		}
		if length > minSize { // 滑动窗口左边=>右移
			window[s[left]]--
			if window[s[left]] == 0 {
				count--
			}
			left++
			length--
		}
		// 只考虑最小值minSize，例如：minSize=2, maxSize=3
		// 如果abc出现3次，那么代表ab至少出现>=3次
		if count <= maxLetters {
			m[s[left:i+1]]++
		}
	}
	for _, v := range m {
		res = max(res, v)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
