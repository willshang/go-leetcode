package main

import "fmt"

func main() {
	fmt.Println(findLUSlength("aba", "cdc"))
}

/*
当两个字符串完全相同，则不存在特殊子序列，返回-1
当两个字符串长度不同，较长字符串组成的序列肯定是特殊序列
当两个字符串长度一样但不完全相同，则任意一个组成的子序列都是特殊序列
*/
// leetcode521_最长特殊序列Ⅰ
func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	}
	return max(len(a), len(b))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
