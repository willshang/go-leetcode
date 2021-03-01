package main

import "sort"

func main() {

}

func longestStrChain(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	res := 0
	dp := make(map[string]int)
	for i := 0; i < len(words); i++ {
		str := words[i]
		for j := 0; j < len(words[i]); j++ {
			target := words[i][:j] + words[i][j+1:]
			if dp[target]+1 > dp[str] {
				dp[str] = dp[target] + 1
			}
		}
		if dp[str] > res {
			res = dp[str]
		}
	}
	return res
}
