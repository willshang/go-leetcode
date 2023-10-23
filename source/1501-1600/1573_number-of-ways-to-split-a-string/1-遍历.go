package main

import "fmt"

func main() {
	fmt.Println(numWays("10101"))
	fmt.Println(numWays("1001"))
	fmt.Println(numWays("0000"))
	fmt.Println(numWays("100100010100110"))
}

func numWays(s string) int {
	total := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			total++
		}
	}
	if total%3 != 0 {
		return 0
	}
	if total == 0 {
		return ((len(s) - 2) * (len(s) - 1) / 2) % 1000000007
	}
	single := total / 3
	first, second := single*1, single*2
	var start, left, right int
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			start++
		} else {
			continue
		}
		if start == first {
			left = i + 1
		} else if start == first+1 {
			left = i - left
		}
		if start == second {
			right = i + 1
		} else if start == second+1 {
			right = i - right
		}
	}
	return (left + 1) * (right + 1) % 1000000007
}
