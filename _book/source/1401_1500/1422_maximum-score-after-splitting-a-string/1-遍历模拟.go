package main

import "fmt"

func main() {
	fmt.Println(maxScore("011101"))
	fmt.Println(maxScore("00"))
}

// leetcode1422_分割字符串的最大得分
func maxScore(s string) int {
	one := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			one++
		}
	}
	max := 0
	zero := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '1' {
			one--
		} else {
			zero++
		}
		if one+zero > max {
			max = one + zero
		}
	}
	return max
}
