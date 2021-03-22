package main

import "fmt"

func main() {
	fmt.Println(countHomogenous("zzzz"))
}

func countHomogenous(s string) int {
	res := 0
	left, right := 0, 0
	for right < len(s) {
		if s[left] != s[right] {
			left = right
		} else {
			res = res + (right-left+1)%1000000007
			right++
		}
	}
	return res % 1000000007
}
