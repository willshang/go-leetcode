package main

import "fmt"

func main() {
	first := TreeNode{Val: 2}
	second := TreeNode{Val: 1}
	third := TreeNode{Val: 3}
	first.Left = &second
	first.Right = &third

	fmt.Println(rangeSumBST(&first, 1, 3))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode938_二叉搜索树的范围和
func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	if root.Val < L {
		return rangeSumBST(root.Right, L, R)
	}
	if root.Val > R {
		return rangeSumBST(root.Left, L, R)
	}
	return root.Val + rangeSumBST(root.Right, L, R) + rangeSumBST(root.Left, L, R)
}
