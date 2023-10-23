package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := dfs(root.Left)
	right := dfs(root.Right)
	if left == right {
		return root
	} else if left > right {
		return lcaDeepestLeaves(root.Left)
	}
	return lcaDeepestLeaves(root.Right)
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left)
	right := dfs(root.Right)
	return 1 + max(left, right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
