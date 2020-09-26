package main

import (
	"fmt"
	"strings"
)

func main() {
	//fmt.Println(stringMatching([]string{"mass", "as", "hero", "superhero"}))
	//fmt.Println(stringMatching([]string{"leetcode", "et", "code"}))
	fmt.Println(stringMatching([]string{"leetcoder", "leetcode", "od", "hamlet", "am"}))
}

// leetcode1408_数组中的字符串匹配
func stringMatching(words []string) []string {
	res := make([]string, 0)
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if i != j && strings.Contains(words[j], words[i]) {
				res = append(res, words[i])
				break
			}
		}
	}
	return res
}
