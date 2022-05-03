package main

import "fmt"

func main() {
	fmt.Println(appealSum("abbca"))
}

// leetcode2262_字符串的总引力
func appealSum(s string) int64 {
	res := int64(0)
	n := len(s)
	dp := make([]int, n+1) // 以长度为i结尾的子字符串总引力
	m := make(map[int]int)
	for i := 0; i < 26; i++ {
		m[i] = -1 // 上一次出现的位置
	}
	for i := 0; i < n; i++ {
		v := int(s[i] - 'a')
		// 1、新字符没出现：在dp[i]基础上加上(i+1)
		// 2、新字符出现过：在dp[i]基础上加上i-m[v](到上一次出现的距离)
		dp[i+1] = dp[i] + (i - m[v])
		m[v] = i
	}
	for i := 1; i <= n; i++ {
		res = res + int64(dp[i])
	}
	return res
}
