package main

import "fmt"

func main() {
	fmt.Println(numDecodings("226"))
}

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	pre := 1
	cur := 1
	for i := 1; i < len(s); i++ {
		temp := cur
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				cur = pre
			} else {
				return 0
			}
		} else if s[i-1] == '1' ||
			(s[i-1] == '2' && s[i] >= '1' && s[i] <= '6') {
			cur = cur + pre
		}
		pre = temp
	}
	return cur
}
