package main

import "fmt"

func main() {
	fmt.Println(verifyPostorder([]int{1, 3, 2, 6, 5}))
}

// 剑指offer_面试题33_二叉搜索树的后序遍历序列
func verifyPostorder(postorder []int) bool {
	return dfs(postorder, 0, len(postorder)-1)
}

func dfs(postorder []int, start, end int) bool {
	if start >= end {
		return true
	}
	i := 0
	for i = start; i < end; i++ {
		if postorder[i] > postorder[end] {
			break
		}
	}
	for j := i + 1; j < end; j++ {
		if postorder[j] < postorder[end] {
			return false
		}
	}
	return dfs(postorder, start, i-1) && dfs(postorder, i, end-1)
}
