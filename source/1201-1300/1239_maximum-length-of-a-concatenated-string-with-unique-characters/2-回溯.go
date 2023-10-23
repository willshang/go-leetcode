package main

import "fmt"

func main() {
	fmt.Println(maxLength([]string{"un", "iq", "ue"}))
}

var res int

func maxLength(arr []string) int {
	res = 0
	dfs(arr, "", 0)
	return res
}

func dfs(arr []string, path string, index int) {
	if check(path) == true && len(path) > res {
		res = len(path)
	}
	for i := index; i < len(arr); i++ {
		dfs(arr, path+arr[i], i+1)
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
