package main

import "fmt"

func main() {
	fmt.Println(minAddToMakeValid("())("))
}

func minAddToMakeValid(S string) int {
	left := 0
	right := 0
	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			left++
		} else {
			if left > 0 {
				left--
			} else {
				right++
			}
		}
	}
	return left + right
}
