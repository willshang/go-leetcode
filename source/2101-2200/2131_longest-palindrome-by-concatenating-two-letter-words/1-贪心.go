package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome([]string{"lc", "cl", "gg"}))
	fmt.Println(longestPalindrome([]string{"ab", "ty", "yt", "lc", "cl", "ab"}))
	fmt.Println(longestPalindrome([]string{"dd", "aa", "bb", "dd", "aa", "dd", "bb", "dd", "aa", "cc", "bb", "cc", "dd", "cc"}))
}

// leetcode2131_连接两字母单词得到的最长回文串
func longestPalindrome(words []string) int {
	res := 0
	m := make(map[string]int)
	for i := 0; i < len(words); i++ {
		m[words[i]]++
	}
	// 相同：偶数次数都可以用上；多个奇数出现次数的只能用1个奇数次数
	// 不同：取反后取最小次数
	mid := false
	for k, v := range m {
		str := string(k[1]) + string(k[0])
		if k == str { // 相同
			if v%2 == 1 { // 奇数个aa可以拿1个放在中间
				mid = true
			}
			res = res + 4*(v/2)
		} else {
			res = res + 2*min(v, m[str]) // 最少的1方
		}
	}
	if mid == true { // 有中间的+2
		res = res + 2
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
