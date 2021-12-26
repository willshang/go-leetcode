package main

import "strings"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode2096_从二叉树一个节点到另一个节点每一步的方向
var a, b []string

func getDirections(root *TreeNode, startValue int, destValue int) string {
	a, b = make([]string, 0), make([]string, 0)
	dfs(root, startValue, destValue, make([]string, 0)) // 构建父节点关系

	i := 0
	for i = 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			break
		}
	}
	return strings.Repeat("U", len(a)-i) + strings.Join(b[i:], "")
}

func dfs(root *TreeNode, startValue, destValue int, path []string) {
	if root.Val == startValue {
		a = make([]string, len(path))
		copy(a, path)
	}
	if root.Val == destValue {
		b = make([]string, len(path))
		copy(b, path)
	}
	if root.Left != nil {
		path = append(path, "L")
		dfs(root.Left, startValue, destValue, path)
		path = path[:len(path)-1]
	}
	if root.Right != nil {
		path = append(path, "R")
		dfs(root.Right, startValue, destValue, path)
		path = path[:len(path)-1]
	}
}
