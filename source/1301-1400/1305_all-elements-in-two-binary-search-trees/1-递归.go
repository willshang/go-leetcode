package main

import "sort"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode1305_两棵二叉搜索树中的所有元素
var res []int

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	res = make([]int, 0)
	dfs(root1)
	dfs(root2)
	sort.Ints(res)
	return res
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	res = append(res, root.Val)
	dfs(root.Left)
	dfs(root.Right)
}
