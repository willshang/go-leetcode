package main

import "fmt"

func main() {
	fmt.Println(checkValidString("()"))
}

func checkValidString(s string) bool {
	maxLeft, minLeft := 0, 0 // 可以有最多left和最少left的数量
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			maxLeft++
			minLeft++
		} else if s[i] == '*' {
			maxLeft++        // *当(用
			if minLeft > 0 { // *当)用
				minLeft--
			}
		} else if s[i] == ')' {
			maxLeft--
			if maxLeft < 0 {
				return false
			}
			if minLeft > 0 {
				minLeft--
			}
		}
	}
	return minLeft == 0
}
