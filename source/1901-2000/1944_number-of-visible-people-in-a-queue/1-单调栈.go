package main

import "fmt"

func main() {
	fmt.Println(canSeePersonsCount([]int{10, 6, 8, 5, 11, 9}))
}

// leetcode1944_队列中可以看到的人数
func canSeePersonsCount(heights []int) []int {
	n := len(heights)
	res := make([]int, n)
	stack := make([]int, 0) // 递减栈
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 {
			res[i]++ // 答案+1，必然存在1个可以看到的人
			if heights[i] > heights[stack[len(stack)-1]] {
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}
		stack = append(stack, i)
	}
	return res
}
