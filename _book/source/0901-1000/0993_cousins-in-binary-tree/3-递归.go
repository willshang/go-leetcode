package main

import "fmt"

func main() {
	first := TreeNode{Val: 5}
	firstLeft := TreeNode{Val: 3}
	firstRight := TreeNode{Val: 7}
	first.Left = &firstLeft
	first.Right = &firstRight
	fmt.Println(isCousins(&first, 3, 7))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode993_二叉树的堂兄弟节点
var fatherM map[int]int
var levelM map[int]int

func isCousins(root *TreeNode, x int, y int) bool {
	fatherM = make(map[int]int)
	levelM = make(map[int]int)
	dfs(root, nil, 0)
	return levelM[x] == levelM[y] && fatherM[x] != fatherM[y]
}

func dfs(root *TreeNode, father *TreeNode, level int) {
	if root == nil {
		return
	}
	if father == nil {
		fatherM[root.Val] = 0
	} else {
		fatherM[root.Val] = father.Val
	}
	levelM[root.Val] = level
	dfs(root.Left, root, level+1)
	dfs(root.Right, root, level+1)
}
