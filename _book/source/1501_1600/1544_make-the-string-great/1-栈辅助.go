package main

import "fmt"

func main() {
	fmt.Println(makeGood("mC"))
}

// leetcode1544_整理字符串
func makeGood(s string) string {
	if len(s) <= 1 {
		return s
	}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		stack = append(stack, s[i])
		if len(stack) > 1 {
			length := len(stack)
			if stack[length-1]-'A'+'a' == stack[length-2] ||
				stack[length-1]-'a'+'A' == stack[length-2] {
				stack = stack[:length-2]
			}
		}
	}
	return string(stack)
}
