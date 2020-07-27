package main

import (
	"fmt"
	"strconv"
)

func main() {
	//fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))
}

// leetcode150_逆波兰表达式求值
func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, v := range tokens {
		length := len(stack)
		if v == "+" || v == "-" || v == "*" || v == "/" {
			a := stack[length-2]
			b := stack[length-1]
			stack = stack[:length-2]
			var value int
			if v == "+" {
				value = a + b
			} else if v == "-" {
				value = a - b
			} else if v == "*" {
				value = a * b
			} else {
				value = a / b
			}
			stack = append(stack, value)
		} else {
			value, _ := strconv.Atoi(v)
			stack = append(stack, value)
		}
	}
	return stack[0]
}
