package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode979_在二叉树中分配硬币
var res int

func distributeCoins(root *TreeNode) int {
	res = 0
	dfs(root)
	return res
}

func dfs(root *TreeNode) int { // 该节点子树多余/需要金币数量
	if root == nil {
		return 0
	}
	left := dfs(root.Left)
	right := dfs(root.Right)
	res = res + abs(left) + abs(right)
	return left + right + root.Val - 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
