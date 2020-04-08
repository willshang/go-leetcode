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

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	cnt := 0
	var dfs func(*TreeNode, int)

	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}

		sum = sum - node.Val
		if sum == 0 {
			cnt++
		}
		dfs(node.Left, sum)
		dfs(node.Right, sum)
	}

	dfs(root, sum)

	return cnt + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}
