package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 将原左子树变为节点的右子树
// 再将原右子树变为当前右子树最右节点的右子树。
func flatten(root *TreeNode) {
	dfs(root, nil)
}

func dfs(root *TreeNode, pre *TreeNode) *TreeNode {
	if root == nil {
		return pre
	}
	pre = dfs(root.Right, pre)
	pre = dfs(root.Left, pre)
	root.Right, root.Left = pre, nil
	pre = root
	return pre
}
