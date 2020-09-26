package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(diffWaysToCompute("2-1-1"))
}

// leetcode241_为运算表达式设计优先级
func diffWaysToCompute(input string) []int {
	if value, err := strconv.Atoi(input); err == nil {
		return []int{value}
	}
	res := make([]int, 0)
	for i := 0; i < len(input); i++ {
		char := string(input[i])
		if char == "+" || char == "-" || char == "*" {
			left := diffWaysToCompute(input[:i])
			right := diffWaysToCompute(input[i+1:])
			for _, leftNum := range left {
				for _, rightNum := range right {
					temp := 0
					if char == "+" {
						temp = leftNum + rightNum
					} else if char == "-" {
						temp = leftNum - rightNum
					} else if char == "*" {
						temp = leftNum * rightNum
					}
					res = append(res, temp)
				}
			}
		}
	}
	return res
}
