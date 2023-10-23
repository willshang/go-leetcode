package main

import "fmt"

func main() {
	fmt.Println(hasAllCodes("00110", 2))
}

// leetcode1461_检查一个字符串是否包含所有长度为K的二进制子串
func hasAllCodes(s string, k int) bool {
	m := make(map[string]bool)
	for i := 0; i <= len(s)-k; i++ {
		m[s[i:i+k]] = true
	}
	return len(m) == 1<<k
}
