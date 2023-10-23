package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode968_监控二叉树
var res int

func minCameraCover(root *TreeNode) int {
	res = 0
	if dfs(root) == 0 { // root未被监控
		res++
	}
	return res
}

// 0=>未被监控， 1=>已监控， 2=>安装监控
func dfs(root *TreeNode) int {
	if root == nil {
		return 1 // 空节点假设有监控
	}
	left, right := dfs(root.Left), dfs(root.Right)
	if left == 1 && right == 1 { // 左右节点都有监控，比如都是空节点，那么当前节点是无未被监控
		return 0
	}
	if left+right >= 3 { // 1+2 / 2+1 / 2+2 => 当前节点已经被监控
		return 1
	}
	res++ // 1+0, 0+1, 0+2, 2+0, 0+0 => 需要安装监控
	return 2
}
