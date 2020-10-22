package main

import "fmt"

func main() {
	fmt.Println(diStringMatch("IDID"))
}

// leetcode942_增减字符串匹配
func diStringMatch(S string) []int {
	res := make([]int, len(S)+1)
	left := 0
	right := len(S)
	for i := 0; i < len(S); i++ {
		if S[i] == 'I' {
			res[i] = left
			left++
		} else {
			res[i] = right
			right--
		}
	}
	res[len(S)] = left
	return res
}
