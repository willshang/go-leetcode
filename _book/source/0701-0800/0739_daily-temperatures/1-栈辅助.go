package main

import "fmt"

func main() {
	fmt.Println(dailyTemperatures([]int{73, 72, 75, 71, 69, 72, 76, 73}))
}

// leetcode739_每日温度
func dailyTemperatures(T []int) []int {
	res := make([]int, len(T))
	stack := make([]int, 0) // 栈保存递减数据的下标
	for i := 0; i < len(T); i++ {
		for len(stack) > 0 && T[i] > T[stack[len(stack)-1]] {
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[last] = i - last
		}
		stack = append(stack, i)
	}
	return res
}
