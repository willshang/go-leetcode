package main

import "fmt"

func main() {
	fmt.Println(minimumTime("1100101"))
	fmt.Println(minimumTime("101"))
}

// leetcode2167_移除所有载有违禁货物车厢所需的最少时间
func minimumTime(s string) int {
	n := len(s)
	res := n
	count := 0
	for i := 0; i < n; i++ {
		if s[i] == '1' {
			count = min(count+2, i+1) // left+middle：左边+中间最小值
		}
		res = min(res, count+n-1-i) // (left+middle)+right
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
