package main

import "strings"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var m map[*TreeNode]*TreeNode // 父节点
var start, dest *TreeNode

func getDirections(root *TreeNode, startValue int, destValue int) string {
	m = make(map[*TreeNode]*TreeNode)
	start, dest = &TreeNode{}, &TreeNode{}
	dfs(root, startValue, destValue)            // 构建父节点关系
	a, b := path(start, root), path(dest, root) // 生成根节点到目标节点的路径
	i := 0
	for i = 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			break
		}
	}
	return strings.Repeat("U", len(a)-i) + strings.Join(b[i:], "")
}

func path(cur *TreeNode, root *TreeNode) []string {
	res := make([]string, 0)
	for cur != root {
		prev := m[cur]
		if cur == prev.Left {
			res = append(res, "L")
		} else {
			res = append(res, "R")
		}
		cur = prev
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}

func dfs(root *TreeNode, startValue, destValue int) {
	if root.Val == startValue {
		start = root
	}
	if root.Val == destValue {
		dest = root
	}
	if root.Left != nil {
		m[root.Left] = root
		dfs(root.Left, startValue, destValue)
	}
	if root.Right != nil {
		m[root.Right] = root
		dfs(root.Right, startValue, destValue)
	}
}
