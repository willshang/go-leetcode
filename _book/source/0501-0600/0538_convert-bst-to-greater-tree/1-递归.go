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

	fmt.Println(convertBST(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode538_把二叉搜索树转换为累加树
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	dfs(root, &sum)
	return root
}

func dfs(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	dfs(root.Right, sum)
	*sum = *sum + root.Val
	root.Val = *sum
	dfs(root.Left, sum)
}
