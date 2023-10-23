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

func helper(node *TreeNode, sum int, path []int, level int) int {
	if node == nil {
		return 0
	}
	res := 0
	if sum == node.Val {
		res = 1
	}
	temp := node.Val
	for i := level - 1; i >= 0; i-- {
		temp = temp + path[i]
		if temp == sum {
			res++
		}
	}
	path[level] = node.Val
	return res + helper(node.Left, sum, path, level+1) +
		helper(node.Right, sum, path, level+1)
}

func pathSum(root *TreeNode, sum int) int {
	return helper(root, sum, make([]int, 1001), 0)
}
