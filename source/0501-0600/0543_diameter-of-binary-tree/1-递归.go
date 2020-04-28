package main

import "fmt"

func main() {
	root := TreeNode{}
	root.Val = 1

	right := TreeNode{}
	right.Val = 5

	rightLeft := TreeNode{}
	rightLeft.Val = 4

	root.Right = &right
	right.Left = &rightLeft

	fmt.Println(diameterOfBinaryTree(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode543_二叉树的直径
var res int

func diameterOfBinaryTree(root *TreeNode) int {
	res = 0
	dfs(root)
	return res
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left)
	right := dfs(root.Right)
	path := max(left, right)
	res = max(left+right, res) // 当前节点最大直径与当前保存最大值比较
	return path + 1            // 以该节点为根的最大深度
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
