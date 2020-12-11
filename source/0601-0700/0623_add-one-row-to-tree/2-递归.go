package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode623_在二叉树中增加一行
func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: v}
	}
	if d == 1 {
		return &TreeNode{Val: v, Left: root}
	}
	dfs(root, v, d)
	return root
}

func dfs(root *TreeNode, v int, d int) {
	if root == nil {
		return
	}
	if d == 2 {
		root.Left = &TreeNode{
			Val:  v,
			Left: root.Left,
		}
		root.Right = &TreeNode{
			Val:   v,
			Right: root.Right,
		}
		return
	}
	dfs(root.Left, v, d-1)
	dfs(root.Right, v, d-1)
}
