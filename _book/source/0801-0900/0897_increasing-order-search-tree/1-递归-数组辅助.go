package main

import "fmt"

func main() {
	first := TreeNode{Val: 5}
	firstLeft := TreeNode{Val: 3}
	firstRight := TreeNode{Val: 7}

	first.Left = &firstLeft
	first.Right = &firstRight

	fmt.Println(increasingBST(&first))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	arr := make([]int, 0)
	dfs(root, &arr)
	if len(arr) == 0 {
		return root
	}
	newRoot := &TreeNode{Val: arr[0]}
	cur := newRoot
	for i := 1; i < len(arr); i++ {
		cur.Right = &TreeNode{Val: arr[i]}
		cur = cur.Right
	}
	return newRoot
}

func dfs(node *TreeNode, arr *[]int) {
	if node == nil {
		return
	}
	dfs(node.Left, arr)
	*arr = append(*arr, node.Val)
	dfs(node.Right, arr)
}
