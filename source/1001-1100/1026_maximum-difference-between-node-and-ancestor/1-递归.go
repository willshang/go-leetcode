package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode1026_节点与其祖先之间的最大差值
var res int

func maxAncestorDiff(root *TreeNode) int {
	res = 0
	if root == nil {
		return 0
	}
	dfs(root, root.Val, root.Val)
	return res
}

func dfs(root *TreeNode, minValue, maxValue int) {
	if root == nil {
		return
	}
	if root.Val > maxValue {
		maxValue = root.Val
	}
	if root.Val < minValue {
		minValue = root.Val
	}
	if maxValue-minValue > res {
		res = maxValue - minValue
	}
	dfs(root.Left, minValue, maxValue)
	dfs(root.Right, minValue, maxValue)
}
