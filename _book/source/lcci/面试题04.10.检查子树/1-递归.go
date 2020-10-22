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

	fmt.Println(checkSubTree(&first, &first1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 程序员面试金典04.10_检查子树
func checkSubTree(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil {
		return false
	}
	return isSame(t1, t2) || checkSubTree(t1.Left, t2) || checkSubTree(t1.Right, t2)
}

func isSame(s *TreeNode, t *TreeNode) bool {
	if s == nil || t == nil {
		return t == s
	}
	return isSame(s.Left, t.Left) && isSame(s.Right, t.Right) && s.Val == t.Val
}
