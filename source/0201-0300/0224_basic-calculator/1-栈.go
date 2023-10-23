package main

import (
	"fmt"
)

func main() {
	//fmt.Println(calculate("1+1"))
	fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)"))
}

// leetcode224_基本计算器
func calculate(s string) int {
	stack := make([]int, 0)
	num := 0
	res := 0
	sign := 1
	for i := 0; i < len(s); i++ {
		if '0' <= s[i] && s[i] <= '9' {
			num = 0
			for ; i < len(s) && '0' <= s[i] && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			res = res + sign*num
			i--
		} else if s[i] == '+' {
			sign = 1
		} else if s[i] == '-' {
			sign = -1
		} else if s[i] == '(' {
			stack = append(stack, res, sign)
			res = 0
			sign = 1
		} else if s[i] == ')' {
			sign = stack[len(stack)-1]
			prev := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			res = prev + sign*res*sign
		}
	}
	return res
}
