package main

func main() {
	root := TreeNode{Val: 2}
	rootfirst := TreeNode{Val: 3}
	rootSecond := TreeNode{Val: 1}
	root.Left = &rootfirst
	root.Right = &rootSecond
	recoverTree(&root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode99_恢复二叉搜索树
var prev, first, second *TreeNode

func recoverTree(root *TreeNode) {
	prev, first, second = nil, nil, nil
	dfs(root)
	first.Val, second.Val = second.Val, first.Val
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	if prev != nil && prev.Val > root.Val {
		second = root
		if first == nil {
			first = prev
		} else {
			return
		}
	}
	prev = root
	dfs(root.Right)
}
