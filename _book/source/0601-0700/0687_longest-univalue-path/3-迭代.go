package main

import "fmt"

func main() {
	root := TreeNode{Val: 5}
	rootLeft := TreeNode{Val: 4}
	rootRight := TreeNode{Val: 5}
	rootLeftLeft := TreeNode{Val: 1}
	rootLeftRight := TreeNode{Val: 1}
	rootRightRight := TreeNode{Val: 5}

	root.Left = &rootLeft
	root.Right = &rootRight

	rootLeft.Left = &rootLeftLeft
	rootLeft.Right = &rootLeftRight
	rootRight.Right = &rootRightRight

	fmt.Println(longestUnivaluePath(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	res := 0
	stack := make([]*TreeNode, 0)
	m := make(map[*TreeNode]int)

	cur := root
	var prev *TreeNode
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		if cur.Right == nil || cur.Right == prev {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			leftLen := 0
			rightLen := 0
			if v, ok := m[cur.Left]; ok {
				leftLen = v
			}
			if v, ok := m[cur.Right]; ok {
				rightLen = v
			}
			var left, right int
			if cur.Left != nil && cur.Val == cur.Left.Val {
				left = leftLen + 1
			}
			if cur.Right != nil && cur.Val == cur.Right.Val {
				right = rightLen + 1
			}

			if left+right > res {
				res = left + right
			}
			if left > right {
				m[cur] = left
			} else {
				m[cur] = right
			}
			prev = cur
			cur = nil
		} else {
			cur = cur.Right
		}
	}
	return res
}
