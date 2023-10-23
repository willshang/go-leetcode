package main

import "fmt"

func main() {
	fmt.Println(calculate("3+2*2"))
}

// leetcode227_基本计算器II
func calculate(s string) int {
	stack := make([]int, 0)
	op := make([]int, 0)
	num := 0
	for i := 0; i < len(s); i++ {
		if '0' <= s[i] && s[i] <= '9' {
			num = 0
			for i < len(s) && '0' <= s[i] && s[i] <= '9' {
				num = num*10 + int(s[i]-'0')
				i++
			}
			// 处理乘除计算
			if len(op) > 0 && op[len(op)-1] > 1 {
				if op[len(op)-1] == 2 {
					stack[len(stack)-1] = stack[len(stack)-1] * num
				} else {
					stack[len(stack)-1] = stack[len(stack)-1] / num
				}
				op = op[:len(op)-1]
			} else {
				stack = append(stack, num)
			}
			i--
		} else if s[i] == '+' {
			op = append(op, 1)
		} else if s[i] == '-' {
			op = append(op, -1)
		} else if s[i] == '*' {
			op = append(op, 2)
		} else if s[i] == '/' {
			op = append(op, 3)
		}
	}
	// 处理加减
	for len(op) > 0 {
		stack[1] = stack[0] + stack[1]*op[0]
		stack = stack[1:]
		op = op[1:]
	}
	return stack[0]
}
