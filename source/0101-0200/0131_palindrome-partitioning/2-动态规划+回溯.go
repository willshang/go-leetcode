package main

import "fmt"

func main() {
	fmt.Println(partition("aab"))
}

var res [][]string
var dp [][]bool

func partition(s string) [][]string {
	res = make([][]string, 0)
	arr := make([]string, 0)
	dp = make([][]bool, len(s))
	for r := 0; r < len(s); r++ {
		dp[r] = make([]bool, len(s))
		dp[r][r] = true
		for l := 0; l < r; l++ {
			if s[l] == s[r] && (r-l <= 2 || dp[l+1][r-1] == true) {
				dp[l][r] = true
			} else {
				dp[l][r] = false
			}
		}
	}
	dfs(s, 0, arr)
	return res
}

func dfs(s string, level int, arr []string) {
	if level == len(s) {
		temp := make([]string, len(arr))
		copy(temp, arr)
		res = append(res, temp)
		return
	}
	for i := level; i < len(s); i++ {
		str := s[level : i+1]
		if dp[level][i] == true {
			dfs(s, i+1, append(arr, str))
		}
	}
}
