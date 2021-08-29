package main

import "fmt"

func main() {
	fmt.Println(minTimeToType("abc"))
}

// leetcode1974_使用特殊打字机键入单词的最少时间
func minTimeToType(word string) int {
	res := 0
	cur := 0
	for i := 0; i < len(word); i++ {
		target := int(word[i] - 'a')
		left := (cur - target + 26) % 26
		right := (target - cur + 26) % 26
		res = res + 1 + min(left, right)
		cur = target
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
