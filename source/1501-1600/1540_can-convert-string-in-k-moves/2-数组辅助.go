package main

import "fmt"

func main() {
	//fmt.Println(canConvertString("input", "ouput", 9))
	//fmt.Println(canConvertString("abc", "bcd", 10))
	fmt.Println(canConvertString("iqssxdlb", "dyuqrwyr", 40))
}

// leetcode1540_K次操作转变字符串
func canConvertString(s string, t string, k int) bool {
	if len(s) != len(t) {
		return false
	}
	next := [26]int{}
	for i := 0; i < 26; i++ {
		next[i] = i
	}
	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			count := (int(t[i]) - int(s[i]) + 26) % 26 // 最少操作的次数
			// 1 <= i <= k，其中i只能出现一次，如果切换的次数一样，需要到下N一圈(i+26xN)才行
			if next[count] > k {
				return false
			}
			next[count] = next[count] + 26
		}
	}
	return true
}
