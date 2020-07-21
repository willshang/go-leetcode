package main

import "fmt"

func main() {
	fmt.Println(partition("aab"))
}

var res [][]string

func partition(s string) [][]string {
	res = make([][]string, 0)
	arr := make([]string, 0)
	dfs(s, 0, arr)
	return res
}

func dfs(s string, level int, arr []string) {
	if level == len(s) {
		temp := make([]string, len(arr))
		copy(temp, arr)
		res = append(res, arr)
		return
	}
	for i := 0; i < len(s)-level; i++ {
		str := s[level : level+i+1]
		if judge(str) == false {
			continue
		}
		dfs(s, level+i+1, append(arr, str))
	}
}

func judge(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
