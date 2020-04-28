package main

import "fmt"

func main() {
	root := TreeNode{}
	root.Val = 1

	right := TreeNode{}
	right.Val = 5

	rightLeft := TreeNode{}
	rightLeft.Val = 4

	root.Right = &right
	right.Left = &rightLeft

	fmt.Println(diameterOfBinaryTree(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := 0
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
			if leftLen > rightLen {
				m[cur] = leftLen + 1
			} else {
				m[cur] = rightLen + 1
			}
			if max < leftLen+rightLen {
				max = leftLen + rightLen
			}
			prev = cur
			cur = nil
		} else {
			cur = cur.Right
		}
	}
	return max
}
