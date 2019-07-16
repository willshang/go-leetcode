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
	preTrav(root, &arr)
	length := len(arr)

	if length == 0 {
		return root
	}
	newroot := &TreeNode{Val: arr[0]}
	cur := newroot

	for i := 1; i < length; i++ {
		cur.Right = &TreeNode{Val: arr[i]}
		cur = cur.Right
	}
	return newroot
}

func preTrav(node *TreeNode, arr *[]int) {
	if node == nil {
		return
	}
	preTrav(node.Left, arr)
	*arr = append(*arr, node.Val)
	preTrav(node.Right, arr)
}
