package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("23"))
}

// leetcode17_电话号码的字母组合
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	arr := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	res := []string{""}
	for i := 0; i < len(digits); i++ {
		length := len(res)
		for j := 0; j < length; j++ {
			for k := 0; k < len(arr[digits[i]-'0']); k++ {
				res = append(res, res[j]+string(arr[digits[i]-'0'][k]))
			}
		}
		res = res[length:]
	}
	return res
}
