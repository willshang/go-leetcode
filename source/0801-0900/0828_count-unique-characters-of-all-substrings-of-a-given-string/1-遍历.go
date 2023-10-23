package main

func main() {

}

// leetcode828_统计子串中的唯一字符
var mod = 1000000007

func uniqueLetterString(s string) int {
	res := 0
	var j, k int
	n := len(s)
	for i := 0; i < n; i++ { // 计算当前字符 作为 唯一字符 的子串的次数
		for j = i - 1; 0 <= j && s[i] != s[j]; j-- {
		}
		for k = i + 1; k < n && s[i] != s[k]; k++ {
		}
		res = (res + (i-j)*(k-i)) % mod // 总数：left * right
	}
	return res
}
