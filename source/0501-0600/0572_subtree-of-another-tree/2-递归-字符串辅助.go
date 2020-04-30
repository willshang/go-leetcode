package main

import (
	"fmt"
	"strings"
)

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

	fmt.Println(isSubtree(&first, &first1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode572_另一个树的子树
func isSubtree(s *TreeNode, t *TreeNode) bool {
	sStr := dfs(s, "")
	tStr := dfs(t, "")
	return strings.Contains(sStr, tStr)
}

func dfs(s *TreeNode, pre string) string {
	if s == nil {
		return pre
	}
	return fmt.Sprintf("#%d%s%s", s.Val, dfs(s.Left, "l"), dfs(s.Right, "r"))
}
