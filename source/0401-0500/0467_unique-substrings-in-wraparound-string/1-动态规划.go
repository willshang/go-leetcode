package main

func main() {

}

// leetcode467_环绕字符串中唯一的子字符串
func findSubstringInWraproundString(p string) int {
	dp := [26]int{}
	count := 0
	for i := 0; i < len(p); i++ {
		value := int(p[i] - 'a')
		if i > 0 && (value-int(p[i-1]-'a')-1)%26 == 0 {
			count++
		} else {
			count = 1
		}
		dp[value] = max(dp[value], count)
	}
	res := 0
	for i := 0; i < len(dp); i++ {
		res = res + dp[i]
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
