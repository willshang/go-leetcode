package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxValue = math.MaxInt32 / 10

func minCameraCover(root *TreeNode) int {
	_, res, _ := dfs(root)
	return res
}

func dfs(root *TreeNode) (a, b, c int) {
	if root == nil {
		return maxValue, 0, 0
	}
	la, lb, lc := dfs(root.Left)
	ra, rb, rc := dfs(root.Right)
	a = lc + rc + 1               // root必须放置摄像头的情况下，覆盖整棵树需要的摄像头数目。
	b = min(a, min(la+rb, ra+lb)) // 覆盖整棵树需要的摄像头数目,无论root是否放置摄像头。
	c = min(a, lb+rb)             // 覆盖两棵子树需要的摄像头数目,无论节点root本身是否被监控到。
	return a, b, c
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
