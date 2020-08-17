package main

import (
	"fmt"
)

func main() {
	first := TreeNode{Val: 1}
	second := TreeNode{Val: 2}
	third := TreeNode{Val: 3}
	third2 := TreeNode{Val: 4}
	third3 := TreeNode{Val: 6}
	first.Left = &second
	first.Right = &third
	second.Left = &third2
	second.Right = &third3

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

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}

	res := false
	list := make([]*TreeNode, 0)
	list = append(list, A)
	for len(list) > 0 {
		node := list[0]
		list = list[1:]
		if node.Val == B.Val {
			res = dfs(node, B)
			if res == true {
				return true
			}
		}
		if node.Left != nil {
			list = append(list, node.Left)
		}
		if node.Right != nil {
			list = append(list, node.Right)
		}
	}
	return res
}

func dfs(A *TreeNode, B *TreeNode) bool {
	fmt.Println(A, B)
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}
	return dfs(A.Left, B.Left) && dfs(A.Right, B.Right) && A.Val == B.Val
}
