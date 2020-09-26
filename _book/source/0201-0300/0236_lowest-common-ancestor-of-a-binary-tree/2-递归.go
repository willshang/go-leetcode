package main

import "fmt"

func main() {
	first := &TreeNode{Val: 6}
	second := &TreeNode{Val: 2}
	third := &TreeNode{Val: 8}
	first.Left = second
	first.Right = third

	fmt.Println(lowestCommonAncestor(first, second, third))
	fmt.Println(first.Val)
	fmt.Println(first.Left.Val)
	fmt.Println(first.Right.Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var m map[int]*TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	m = make(map[int]*TreeNode)
	dfs(root)
	visited := make(map[int]bool)
	for p != nil {
		visited[p.Val] = true
		p = m[p.Val]
	}
	for q != nil {
		if visited[q.Val] == true {
			return q
		}
		q = m[q.Val]
	}
	return nil
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		m[root.Left.Val] = root
		dfs(root.Left)
	}
	if root.Right != nil {
		m[root.Right.Val] = root
		dfs(root.Right)
	}
}
