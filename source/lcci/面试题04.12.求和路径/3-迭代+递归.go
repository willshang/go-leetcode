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
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	res := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		tempSum := 0
		res += dfs(node, sum, tempSum)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return res
}

func dfs(node *TreeNode, sum int, curSum int) int {
	res := 0
	curSum = curSum + node.Val
	if curSum == sum {
		res++
	}
	if node.Left != nil {
		res += dfs(node.Left, sum, curSum)
	}
	if node.Right != nil {
		res += dfs(node.Right, sum, curSum)
	}
	return res
}
