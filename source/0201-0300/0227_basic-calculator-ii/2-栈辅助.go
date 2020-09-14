package main

import (
	"fmt"
	"strings"
)

func main() {
	//fmt.Println(calculate("3+2*2"))
	fmt.Println(calculate("3/2 "))
}

func calculate(s string) int {
	s = strings.Trim(s, " ") // 避免"3/2 "的情况
	stack := make([]int, 0)
	num := 0
	sign := byte('+')
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		if '0' <= s[i] && s[i] <= '9' {
			num = num*10 + int(s[i]-'0')
		}
		if s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i] == '/' || i == len(s)-1 {
			// 处理前一个符号
			switch sign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				prev := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, num*prev)
			case '/':
				prev := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, prev/num)
			}
			num = 0
			sign = s[i]
		}
	}
	res := 0
	for i := 0; i < len(stack); i++ {
		res = res + stack[i]
	}
	return res
}
