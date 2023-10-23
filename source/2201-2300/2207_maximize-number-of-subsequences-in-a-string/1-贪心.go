package main

import "strings"

func main() {

}

// leetcode2207_字符串中最多数目的子字符串
func maximumSubsequenceCount(text string, pattern string) int64 {
	a, b := pattern[0], pattern[1]
	n := len(text)
	if a == b { // 相同情况：count+1的组合
		count := int64(strings.Count(text, string(a)))
		return (count + 1) * count / 2
	}
	var countA, countB int64
	res := int64(0)
	for i := 0; i < n; i++ {
		if text[i] == a {
			countA++
		} else if text[i] == b {
			countB++
			res = res + countA // 子序列数量：加上前面的a
		}
	}
	return res + max(countA, countB) // 最后决定插入a还是b：插入a在最前面，插入b在最后面
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
