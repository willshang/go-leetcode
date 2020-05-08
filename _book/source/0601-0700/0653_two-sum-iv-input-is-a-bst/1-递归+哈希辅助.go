package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode653_两数之和IV-输入BST
func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	m := map[int]int{}
	return dfs(root, k, m)
}

func dfs(node *TreeNode, k int, m map[int]int) bool {
	if node == nil {
		return false
	}
	if _, ok := m[k-node.Val]; ok {
		return true
	}
	m[node.Val] = node.Val
	return dfs(node.Left, k, m) || dfs(node.Right, k, m)
}
