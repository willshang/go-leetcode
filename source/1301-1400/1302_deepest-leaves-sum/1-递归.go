package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode1302_层数最深叶子节点的和
var maxLevel, sum int

func deepestLeavesSum(root *TreeNode) int {
	maxLevel, sum = 0, 0
	dfs(root, 0)
	return sum
}

func dfs(root *TreeNode, level int) {
	if root != nil {
		if level > maxLevel {
			maxLevel = level
			sum = root.Val
		} else if level == maxLevel {
			sum = sum + root.Val
		}
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}
}
