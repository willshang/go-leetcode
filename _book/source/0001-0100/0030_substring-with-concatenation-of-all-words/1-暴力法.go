package main

import "fmt"

func main() {
	fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
}

// leetcode30_串联所有单词的子串
func findSubstring(s string, words []string) []int {
	res := make([]int, 0)
	length, n := len(s), len(words)
	if length == 0 || n == 0 || len(words[0]) == 0 {
		return res
	}
	single := len(words[0])
	m := make(map[string]int)
	for i := 0; i < len(words); i++ {
		m[words[i]]++
	}
	for i := 0; i <= length-n*single; i++ {
		temp := make(map[string]int)
		for j := 0; j < n; j++ {
			l := i + j*single
			str := s[l : l+single]
			if _, ok := m[str]; !ok {
				break
			}
			temp[str]++
			if temp[str] > m[str] {
				break
			}
			if compare(m, temp) == true {
				res = append(res, i)
				break
			}
		}
	}
	return res
}

func compare(m1, m2 map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	return true
}
