package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int
var maxLevel int

func findBottomLeftValue(root *TreeNode) int {
	res = 0
	maxLevel = -1
	if root == nil {
		return res
	}
	dfs(root, 0)
	return res
}

func dfs(root *TreeNode, level int) {
	if root == nil {
		return
	}
	dfs(root.Left, level+1)
	if level > maxLevel {
		maxLevel = level
		res = root.Val
	}
	dfs(root.Right, level+1)
}
