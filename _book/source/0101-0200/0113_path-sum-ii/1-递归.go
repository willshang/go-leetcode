package main

import "fmt"

func main() {
	first := TreeNode{Val: 1}
	second := TreeNode{Val: 2}
	third := TreeNode{Val: 3}
	first.Left = &second
	first.Right = &third
	fmt.Println(pathSum(&first, 4))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode113_路径总和II
var res [][]int

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	res = make([][]int, 0)
	var arr []int
	dfs(root, sum, arr)
	return res
}

func dfs(root *TreeNode, sum int, arr []int) {
	if root == nil {
		return
	}
	arr = append(arr, root.Val)
	if root.Val == sum && root.Left == nil && root.Right == nil {
		temp := make([]int, len(arr))
		copy(temp, arr)
		res = append(res, temp)
	}
	dfs(root.Left, sum-root.Val, arr)
	dfs(root.Right, sum-root.Val, arr)
	arr = arr[:len(arr)-1]
}
