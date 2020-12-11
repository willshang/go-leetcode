package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode1161_最大层内元素和
var arr [][]int

func maxLevelSum(root *TreeNode) int {
	arr = make([][]int, 0)
	if root == nil {
		return 0
	}
	dfs(root, 0)
	res := 0
	maxValue := math.MinInt32
	for i := 0; i < len(arr); i++ {
		sum := 0
		for j := 0; j < len(arr[i]); j++ {
			sum = sum + arr[i][j]
		}
		if sum > maxValue {
			maxValue = sum
			res = i + 1
		}
	}
	return res
}

func dfs(root *TreeNode, level int) {
	if root == nil {
		return
	}
	if level == len(arr) {
		arr = append(arr, []int{})
	}
	arr[level] = append(arr[level], root.Val)
	dfs(root.Left, level+1)
	dfs(root.Right, level+1)
}
