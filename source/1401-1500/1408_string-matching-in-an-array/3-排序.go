package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	//fmt.Println(stringMatching([]string{"mass", "as", "hero", "superhero"}))
	//fmt.Println(stringMatching([]string{"leetcode", "et", "code"}))
	fmt.Println(stringMatching([]string{"leetcoder", "leetcode", "od", "hamlet", "am"}))
}

func stringMatching(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	res := make([]string, 0)
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if strings.Contains(words[j], words[i]) {
				res = append(res, words[i])
				break
			}
		}
	}
	return res
}
