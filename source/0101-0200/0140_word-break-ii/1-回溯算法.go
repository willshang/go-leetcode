package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
}

var res []string
var m map[string]bool

func wordBreak(s string, wordDict []string) []string {
	m = make(map[string]bool)
	for _, word := range wordDict {
		m[word] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] == true && m[s[j:i]] == true {
				dp[i] = true
				break
			}
		}
	}
	if dp[len(s)] == false {
		return nil
	}
	res = make([]string, 0)
	dfs(s, make([]string, 0))
	return res
}

func dfs(str string, arr []string) {
	if len(str) == 0 {
		res = append(res, strings.Join(arr, " "))
		return
	}
	for i := 1; i <= len(str); i++ {
		if m[str[:i]] == true {
			dfs(str[i:], append(arr, str[:i]))
		}
	}
}
