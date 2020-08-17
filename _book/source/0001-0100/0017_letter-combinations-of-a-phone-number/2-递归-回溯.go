package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("23"))
}

var res []string
var arr = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	res = make([]string, 0)
	dfs(digits, 0, "")
	return res
}

func dfs(digits string, index int, str string) {
	if index == len(digits) {
		res = append(res, str)
		return
	}
	for i := 0; i < len(arr[digits[index]-'0']); i++ {
		dfs(digits, index+1, str+string(arr[digits[index]-'0'][i]))
	}
}
