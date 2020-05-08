package main

import "fmt"

func main() {
	first := TreeNode{Val: 1}
	second := TreeNode{Val: 2}
	third := TreeNode{Val: 0}

	first.Left = &third
	first.Right = &second

	fmt.Println(trimBST(&first, 1, 2))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return nil
	}
	// 找到根节点
	for root.Val < L || root.Val > R {
		if root.Val < L {
			root = root.Right
		} else {
			root = root.Left
		}
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	cur := root
	temp := root
	for len(stack) > 0 {
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur.Left != nil {
			if cur.Left.Val >= L {
				// 左节点>=L，继续向左
				stack = append(stack, cur.Left)
			} else {
				// 在当前左节点，向它的右节点找到满足<L的值，
				// 并把当前左指针指向找到的值
				// 如示例2里面的，3的Left指向找到的2
				// 然后入栈继续在2找
				temp = cur.Left
				for temp != nil && temp.Val < L {
					temp = temp.Right
				}
				cur.Left = temp
				if temp != nil {
					stack = append(stack, temp)
				}
			}
		}
		if cur.Right != nil {
			if cur.Right.Val <= R {
				stack = append(stack, cur.Right)
			} else {
				temp = cur.Right
				for temp != nil && temp.Val > R {
					temp = temp.Left
				}
				cur.Right = temp
				if temp != nil {
					stack = append(stack, temp)
				}
			}
		}
	}
	return root
}
