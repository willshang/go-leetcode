package main

import "fmt"

func main() {
	fmt.Println(maxChunksToSorted([]int{3, 1, 2, 4, 4}))
}

func maxChunksToSorted(arr []int) int {
	stack := make([]int, 0) // 递增栈
	for i := 0; i < len(arr); i++ {
		if len(stack) > 0 && arr[i] < stack[len(stack)-1] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			for len(stack) > 0 && arr[i] < stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, top)
		} else {
			stack = append(stack, arr[i])
		}
	}
	return len(stack)
}
