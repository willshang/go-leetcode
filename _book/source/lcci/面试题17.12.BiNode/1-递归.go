package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 程序员面试金典17.12_BiNode
func convertBiNode(root *TreeNode) *TreeNode {
	head := &TreeNode{}
	cur := head
	dfs(root, cur)
	return head.Right
}

func dfs(root, cur *TreeNode) *TreeNode {
	if root != nil {
		cur = dfs(root.Left, cur)
		root.Left = nil
		cur.Right = root
		cur = root
		cur = dfs(root.Right, cur)
	}
	return cur
}
