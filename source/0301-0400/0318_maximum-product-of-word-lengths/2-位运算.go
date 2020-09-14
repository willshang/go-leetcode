package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxProduct([]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}))
}

// leetcode318_最大单词长度乘积
func maxProduct(words []string) int {
	res := 0
	arr := make([]int, len(words))
	for i := 0; i < len(words); i++ {
		for _, char := range words[i] {
			// 位或 只要有1，那么就是1
			arr[i] = arr[i] | 1<<uint(char-'a')
		}
	}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]&arr[j] == 0 && res < len(words[i])*len(words[j]) {
				res = len(words[i]) * len(words[j])
			}
		}
	}
	return res
}
