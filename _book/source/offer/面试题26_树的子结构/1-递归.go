package main

import "fmt"

func main() {
	first := TreeNode{Val: 1}
	second := TreeNode{Val: 2}
	third := TreeNode{Val: 3}
	first.Left = &second
	first.Right = &third

	first1 := TreeNode{Val: 1}
	second1 := TreeNode{Val: 2}
	third1 := TreeNode{Val: 3}
	first1.Left = &second1
	first1.Right = &third1

	fmt.Println(isSubStructure(&first, &first1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 剑指offer_面试题26_树的子结构
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	return dfs(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func dfs(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}
	return dfs(A.Left, B.Left) && dfs(A.Right, B.Right) && A.Val == B.Val
}
