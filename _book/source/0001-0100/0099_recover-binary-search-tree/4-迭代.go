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

func recoverTree(root *TreeNode) {
	var prev, temp, first, second *TreeNode
	for root != nil {
		temp = root.Left
		if temp != nil {
			// 当前root节点向左走一步，然后一直向右走至无法走为止
			for temp.Right != nil && temp.Right != root {
				temp = temp.Right
			}
			if temp.Right == nil {
				temp.Right = root
				root = root.Left
				continue
			} else {
				temp.Right = nil
			}
		}
		if prev != nil && prev.Val > root.Val {
			second = root
			if first == nil {
				first = prev
			}
		}
		prev = root
		root = root.Right
	}
	first.Val, second.Val = second.Val, first.Val
}
