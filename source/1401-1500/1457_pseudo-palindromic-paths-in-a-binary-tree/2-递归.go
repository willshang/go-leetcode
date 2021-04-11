package main

import "fmt"

func main() {
	fmt.Println()
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int

func pseudoPalindromicPaths(root *TreeNode) int {
	res = 0
	dfs(root, make([]int, 10))
	return res
}

func dfs(root *TreeNode, arr []int) {
	if root == nil {
		return
	}
	arr[root.Val]++
	if root.Left == nil && root.Right == nil {
		count := 0
		for i := 1; i <= 9; i++ {
			if arr[i]%2 == 1 {
				count++
			}
		}
		if count <= 1 {
			res++
		}
		return
	}
	tempLeft := make([]int, 10)
	copy(tempLeft, arr)
	dfs(root.Left, tempLeft)

	tempRight := make([]int, 10)
	copy(tempRight, arr)
	dfs(root.Right, tempRight)
}
