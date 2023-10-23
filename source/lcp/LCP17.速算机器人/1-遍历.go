package main

import "fmt"

func main() {
	fmt.Println(calculate("AB"))
}

// leetcode_lcp17_速算机器人
func calculate(s string) int {
	x, y := 1, 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			x = 2*x + y
		} else if s[i] == 'B' {
			y = 2*y + x
		}
	}
	return x + y
}
