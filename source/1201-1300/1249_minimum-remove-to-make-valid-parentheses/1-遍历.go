package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(minRemoveToMakeValid("lee(t(c)o)de)"))
}

func minRemoveToMakeValid(s string) string {
	arr := []byte(s)
	sum := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == '(' {
			sum++
		} else if arr[i] == ')' {
			sum--
			if sum < 0 {
				arr[i] = ' '
				sum = 0
			}
		}
	}
	sum = 0
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == ')' {
			sum++
		} else if arr[i] == '(' {
			sum--
			if sum < 0 {
				arr[i] = ' '
				sum = 0
			}
		}
	}
	return strings.ReplaceAll(string(arr), " ", "")
}
