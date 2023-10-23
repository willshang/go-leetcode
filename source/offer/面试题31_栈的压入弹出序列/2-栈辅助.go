package main

import "fmt"

func main() {
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
}

func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0)
	res := false
	i := 0
	j := 0
	for j < len(popped) {
		for len(stack) == 0 || stack[len(stack)-1] != popped[j] {
			if i == len(pushed) {
				break
			}
			stack = append(stack, pushed[i])
			i++
		}
		if stack[len(stack)-1] != popped[j] {
			break
		}
		stack = stack[:len(stack)-1]
		j++
	}
	if len(stack) == 0 && j == len(popped) {
		res = true
	}
	return res
}
