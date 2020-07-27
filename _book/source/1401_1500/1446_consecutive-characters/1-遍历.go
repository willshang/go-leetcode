package main

import "fmt"

func main() {
	fmt.Println(maxPower("leetcode"))
}

// leetcode1446_连续字符
func maxPower(s string) int {
	max := 1
	count := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			count = 1
		}
		if count > max {
			max = count
		}
	}
	return max
}
