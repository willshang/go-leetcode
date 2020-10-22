package main

import "fmt"

func main() {
	root := TreeNode{Val: 1}
	rootfirst := TreeNode{Val: 2}
	rootSecond := TreeNode{Val: 3}
	root.Left = &rootfirst
	root.Right = &rootSecond
	fmt.Println(levelOrder(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

var res []*ListNode

func listOfDepth(tree *TreeNode) []*ListNode {
	level := 0
	res = make([]*ListNode, 0)
	dfs(tree, level)
	return res
}

func dfs(root *TreeNode, level int) {
	if root == nil {
		return
	}
	if level >= len(res) {
		res = append(res, &ListNode{root.Val, nil})
	} else {
		head := res[level]
		for head.Next != nil {
			head = head.Next
		}
		head.Next = &ListNode{root.Val, nil}
	}
	dfs(root.Left, level+1)
	dfs(root.Right, level+1)
}
