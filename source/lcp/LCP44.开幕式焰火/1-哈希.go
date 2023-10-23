package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode_lcp44_开幕式焰火
var m map[int]int

func numColor(root *TreeNode) int {
	m = make(map[int]int)
	dfs(root)
	return len(m)
}

func dfs(root *TreeNode) {
	if root != nil {
		m[root.Val] = 1
		dfs(root.Left)
		dfs(root.Right)
	}
}
