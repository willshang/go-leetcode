package main

import "fmt"

func main() {
	fmt.Println(countHomogenous("zzzz"))
}

// leetcode1759_统计同构子字符串的数目
func countHomogenous(s string) int {
	res := 1
	count := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			count = 1
		}
		res = res + count
	}
	return res % 1000000007
}
