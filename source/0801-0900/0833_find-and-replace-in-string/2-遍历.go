package main

import (
	"fmt"
)

func main() {
	a := "jjievdtjfb"
	arr := []int{4, 6, 1}
	a1 := []string{"md", "tjgb", "jf"}
	a2 := []string{"foe", "oov", "e"}
	fmt.Println(findReplaceString(a, arr, a1, a2))
}

// leetcode833_字符串中的查找与替换
func findReplaceString(S string, indexes []int, sources []string, targets []string) string {
	m := make(map[int]int)
	for i := 0; i < len(indexes); i++ {
		if S[indexes[i]:indexes[i]+len(sources[i])] == sources[i] {
			m[indexes[i]] = i
		}
	}
	res := make([]byte, 0)
	for i := 0; i < len(S); {
		if v, ok := m[i]; ok {
			res = append(res, targets[v]...)
			i = i + len(sources[v])
		} else {
			res = append(res, S[i])
			i++
		}
	}
	return string(res)
}
