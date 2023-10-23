package main

import "fmt"

func main() {
	fmt.Println(numMatchingSubseq("abcde", []string{"ace"}))
}

// leetcode792_匹配子序列的单词数
func numMatchingSubseq(S string, words []string) int {
	res := 0
	m := make(map[string]int)
	for i := 0; i < len(words); i++ {
		m[words[i]]++
	}
	for k, v := range m {
		if judge(S, k) == true {
			res = res + v
		}
	}
	return res
}

func judge(S string, str string) bool {
	for i, j := 0, 0; i < len(S) && j < len(str); i++ {
		if S[i] == str[j] {
			j++
		}
		if j == len(str) {
			return true
		}
	}
	return false
}
