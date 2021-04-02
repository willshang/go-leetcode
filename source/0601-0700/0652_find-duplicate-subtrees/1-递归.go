package main

import "strconv"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode652_寻找重复的子树
var m map[string]int
var res []*TreeNode

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	m = make(map[string]int)
	res = make([]*TreeNode, 0)

	dfs(root)
	return res
}

func dfs(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	value := strconv.Itoa(root.Val) + "," + dfs(root.Left) + "," + dfs(root.Right)
	m[value]++
	if m[value] == 2 {
		res = append(res, root)
	}
	return value
}
