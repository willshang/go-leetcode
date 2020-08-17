package main

import "fmt"

func main() {
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
}

// 剑指offer_面试题31_栈的压入弹出序列
func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0)
	j := 0
	for i := 0; i < len(pushed); i++ {
		stack = append(stack, pushed[i])
		for len(stack) > 0 && stack[len(stack)-1] == popped[j] {
			stack = stack[:len(stack)-1]
			j++
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
