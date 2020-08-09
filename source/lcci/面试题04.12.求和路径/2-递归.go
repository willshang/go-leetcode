package main

import "fmt"

func main() {
	root := TreeNode{Val: 1}
	rootfirst := TreeNode{Val: 2}
	rootSecond := TreeNode{Val: 3}
	root.Left = &rootfirst
	root.Right = &rootSecond

	secondfirst := TreeNode{Val: 3}
	secondSecond := TreeNode{Val: 2}
	rootfirst.Left = &secondfirst
	rootSecond.Right = &secondSecond

	fmt.Println(pathSum(&root, 5))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, sum int) int {
	if node == nil {
		return 0
	}
	sum = sum - node.Val
	res := 0
	if sum == 0 {
		res = 1
	}
	return res + dfs(node.Left, sum) + dfs(node.Right, sum)
}

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return dfs(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}
