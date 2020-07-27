package main

import "fmt"

func main() {
	fmt.Println(numSub("0110111"))
}

// leetcode1513_仅含1的子串数
func numSub(s string) int {
	res := 0
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			count++
		} else {
			count = 0
		}
		res = (res + count) % 1000000007
	}
	return res
}
