package main

import "fmt"

func main() {
	fmt.Println(maxFreq("aaaa", 1, 3, 3))
}

// leetcode1297_子串的最大出现次数
func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	res := 0
	m := make(map[string]int)
	for i := 0; i <= len(s)-minSize; i++ {
		str := s[i : i+minSize]
		count := getCount(str)
		// 只考虑最小值minSize，例如：minSize=2, maxSize=3
		// 如果abc出现3次，那么代表ab至少出现>=3次
		if count <= maxLetters {
			m[str]++
		}
	}
	for _, v := range m {
		res = max(res, v)
	}
	return res
}

func getCount(str string) int {
	m := make(map[byte]int)
	for i := 0; i < len(str); i++ {
		m[str[i]]++
	}
	return len(m)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
