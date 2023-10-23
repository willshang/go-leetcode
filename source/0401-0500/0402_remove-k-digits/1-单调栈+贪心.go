package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(removeKdigits("1432219", 3))
}

// leetcode402_移掉K位数字
func removeKdigits(num string, k int) string {
	stack := make([]byte, 0)
	res := ""
	for i := 0; i < len(num); i++ {
		value := num[i]
		// 栈顶元素打大于后面的元素，摘除栈顶元素（因为前面的更大，需要删除了才能变的最小）
		for len(stack) > 0 && stack[len(stack)-1] > value && k > 0 {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, value)
	}
	stack = stack[:len(stack)-k]
	res = strings.TrimLeft(string(stack), "0")
	if res == "" {
		return "0"
	}
	return res
}
