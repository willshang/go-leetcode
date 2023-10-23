package main

import "fmt"

func main() {
	root := TreeNode{Val: 1}
	rootfirst := TreeNode{Val: 2}
	rootSecond := TreeNode{Val: 3}
	root.Left = &rootfirst
	root.Right = &rootSecond
	fmt.Println(levelOrder(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res [][]int

func levelOrder(root *TreeNode) [][]int {
	res = make([][]int, 0)
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
	if level == len(res) {
		res = append(res, []int{})
	}
	res[level] = append(res[level], root.Val)
	dfs(root.Left, level+1)
	dfs(root.Right, level+1)
}
