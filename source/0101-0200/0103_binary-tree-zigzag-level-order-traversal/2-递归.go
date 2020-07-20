package main

import "fmt"

func main() {
	root := TreeNode{Val: 1}
	rootfirst := TreeNode{Val: 2}
	rootSecond := TreeNode{Val: 3}
	root.Left = &rootfirst
	root.Right = &rootSecond
	fmt.Println(zigzagLevelOrder(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res [][]int

func zigzagLevelOrder(root *TreeNode) [][]int {
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
	if level%2 == 1 {
		arr := res[level]
		arr = append([]int{root.Val}, arr...)
		res[level] = arr
	} else {
		res[level] = append(res[level], root.Val)
	}
	dfs(root.Left, level+1)
	dfs(root.Right, level+1)
}
