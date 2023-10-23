package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode1038_把二叉搜索树转换为累加树
func bstToGst(root *TreeNode) *TreeNode {
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
