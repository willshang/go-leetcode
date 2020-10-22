package main

import "fmt"

func main() {
	fmt.Println(removePalindromeSub("ababa"))
}

// leetcode1332_删除回文子序列
/*
1.长度为0，返回0
2.字符串为回文子序列，返回1
3.字符串不为回文子序列，返回2，因为可以把a或者b一次都去除，题目没有要求去除的是连续的
*/
func removePalindromeSub(s string) int {
	if len(s) == 0 {
		return 0
	}
	for i, j := 0, len(s)-1; i < j; {
		if s[i] != s[j] {
			return 2
		}
		i++
		j--
	}
	return 1
}
