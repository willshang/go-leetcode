package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(verifyPostorder([]int{1, 3, 2, 6, 5}))
	fmt.Println(verifyPostorder([]int{5, 7, 6, 9, 11, 10, 8}))
}

func verifyPostorder(postorder []int) bool {
	if len(postorder) <= 2 {
		return true
	}
	stack := make([]int, 0)
	rootValue := math.MaxInt32
	for i := len(postorder) - 1; i >= 0; i-- {
		if postorder[i] > rootValue {
			// 左子树元素必须要小于递增栈根节点
			return false
		}
		// 数组元素小于单调栈的元素了，表示往左子树走了，记录上个根节点
		for len(stack) > 0 && postorder[i] < stack[len(stack)-1] {
			rootValue = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, postorder[i])
	}
	return true
}
