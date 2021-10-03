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

func helper(node *TreeNode, targetSum int, curSum int) int {
	res := 0
	curSum = curSum + node.Val
	if curSum == targetSum {
		res++
	}
	if node.Left != nil {
		res += helper(node.Left, targetSum, curSum)
	}
	if node.Right != nil {
		res += helper(node.Right, targetSum, curSum)
	}
	return res
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
		res += helper(node, sum, tempSum)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return res
}
