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

// 剑指OfferII050.向下的路径节点之和
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := 0
	var helper func(*TreeNode, int)
	helper = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}
		sum = sum - node.Val
		// 路径不需要从根节点开始，也不需要在叶子节点结束
		if sum == 0 {
			res++
		}
		helper(node.Left, sum)
		helper(node.Right, sum)
	}
	helper(root, targetSum)
	return res + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
}
