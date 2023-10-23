package main

import (
	"fmt"
)

func main() {
	fmt.Println(addOperators("123", 6))
}

var res []string

func addOperators(num string, target int) []string {
	res = make([]string, 0)
	dfs(num, target, 0, "")
	return res
}

func dfs(num string, target int, index int, str string) {
	if index == len(num) {
		// 全排列再计算
		if calculate(str) == target {
			res = append(res, str)
		}
		return
	}
	for i := index; i < len(num); i++ {
		if num[index] == '0' && index < i { // 105 5 => 1*05不符合要求
			return
		}
		s := num[index : i+1]
		if index == 0 {
			dfs(num, target, i+1, str+s)
		} else {
			dfs(num, target, i+1, str+"+"+s)
			dfs(num, target, i+1, str+"-"+s)
			dfs(num, target, i+1, str+"*"+s)
		}
	}
}

// leetcode227.基本计算器II
func calculate(s string) int {
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
