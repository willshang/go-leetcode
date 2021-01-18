package main

import "fmt"

func main() {
	fmt.Println(maxLength([]string{"un", "iq", "ue"}))
}

// leetcode1239_串联字符串的最大长度
var res int

func maxLength(arr []string) int {
	res = 0
	dfs(arr, "", 0)
	return res
}

func dfs(arr []string, path string, index int) {
	for i := index; i < len(arr); i++ {
		newStr := path + arr[i]
		if check(newStr) == true {
			if len(newStr) > res {
				res = len(newStr)
			}
			dfs(arr, path+arr[i], i+1)
		}
	}
}

func check(s string) bool {
	arr := [26]int{}
	for i := 0; i < len(s); i++ {
		value := int(s[i] - 'a')
		arr[value]++
		if arr[value] > 1 {
			return false
		}
	}
	return true
}
