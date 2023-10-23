package main

import "fmt"

func main() {
	fmt.Println(maxPower("leetcode"))
}

func maxPower(s string) int {
	max := 1
	left := 0
	right := 1
	for right < len(s) {
		if s[left] != s[right] {
			if right-left > max {
				max = right - left
			}
			left = right
		}
		right++
	}
	if right-left > max {
		return right - left
	}
	return max
}
