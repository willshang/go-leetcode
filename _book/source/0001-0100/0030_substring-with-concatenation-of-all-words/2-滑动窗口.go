package main

import "fmt"

func main() {
	fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
}

func findSubstring(s string, words []string) []int {
	res := make([]int, 0)
	length := len(s)
	n := len(words)
	if length == 0 || n == 0 || len(words[0]) == 0 {
		return res
	}
	single := len(words[0])
	m := make(map[string]int)
	for i := 0; i < len(words); i++ {
		m[words[i]]++
	}
	for i := 0; i < single; i++ {
		left, right, count := i, i, 0
		temp := make(map[string]int)
		for right+single <= length {
			str := s[right : right+single]
			right = right + single
			if m[str] > 0 {
				temp[str]++
				if temp[str] == m[str] {
					count++
				}
			}
			if right-left == n*single {
				if count == len(m) {
					res = append(res, left)
				}
				leftStr := s[left : left+single]
				left = left + single
				if m[leftStr] > 0 {
					if temp[leftStr] == m[leftStr] {
						count--
					}
					temp[leftStr]--
				}
			}
		}
	}
	return res
}
