package main

func main() {

}

// leetcode2002_两个回文子序列长度的最大乘积
var res int

func maxProduct(s string) int {
	res = 0
	dfs(s, "", "", 0)
	return res
}

func dfs(s string, a, b string, index int) {
	if len(a) > 0 && len(b) > 0 &&
		isPalindrome(a, 0, len(a)-1) && isPalindrome(b, 0, len(b)-1) {
		res = max(res, len(a)*len(b))
	}
	if index == len(s) {
		return
	}
	dfs(s, a, b, index+1)                  // a,b都不选
	dfs(s, a+string(s[index]), b, index+1) // a不选
	dfs(s, a, b+string(s[index]), index+1) // b不选
}

func isPalindrome(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
