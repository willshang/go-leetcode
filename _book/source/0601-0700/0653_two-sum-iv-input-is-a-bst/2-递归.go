package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	return dfs(root, root, k)
}

func dfs(root, searchRoot *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	found := findNode(searchRoot, k-root.Val)
	if found != nil && found != root {
		return true
	}
	return dfs(root.Left, searchRoot, k) ||
		dfs(root.Right, searchRoot, k)
}

func findNode(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == target {
		return root
	}
	if root.Val < target {
		return findNode(root.Right, target)
	}
	return findNode(root.Left, target)
}
