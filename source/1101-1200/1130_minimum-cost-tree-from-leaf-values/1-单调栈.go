package main

import "math"

func main() {

}

func mctFromLeafValues(arr []int) int {
	res := 0
	stack := make([]int, 0) // 单调递减栈
	stack = append(stack, math.MaxInt32)
	for i := 0; i < len(arr); i++ {
		for len(stack) > 0 && arr[i] >= stack[len(stack)-1] { // 大于栈顶
			middle := stack[len(stack)-1] // 中间
			stack = stack[:len(stack)-1]
			left := stack[len(stack)-1]
			right := arr[i]
			res = res + middle*min(left, right)
		}
		stack = append(stack, arr[i])
	}
	for len(stack) > 2 {
		a := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		b := stack[len(stack)-1]
		res = res + a*b
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
