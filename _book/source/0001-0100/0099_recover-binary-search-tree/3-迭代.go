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
	var prev, first, second *TreeNode
	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if prev != nil && root.Val < prev.Val {
			second = root
			if first == nil {
				first = prev
			} else {
				break
			}
		}
		prev = root
		root = root.Right
	}
	first.Val, second.Val = second.Val, first.Val
}
