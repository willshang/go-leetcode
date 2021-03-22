package main

import "fmt"

func main() {
	fmt.Println(countHomogenous("zzzz"))
}

func countHomogenous(s string) int {
	res := 0
	left, right := 0, 0
	for right < len(s) {
		if s[left] == s[right] {
			right++
		} else {
			length := right - left
			res = res + length*(length+1)/2
			left = right
			right++
		}
	}
	length := right - left
	res = res + length*(length+1)/2
	return res % 1000000007
}
