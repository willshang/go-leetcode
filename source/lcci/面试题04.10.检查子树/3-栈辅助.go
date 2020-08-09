package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	first := TreeNode{Val: 1}
	second := TreeNode{Val: 2}
	third := TreeNode{Val: 3}
	first.Left = &second
	first.Right = &third

	first1 := TreeNode{Val: 1}
	second1 := TreeNode{Val: 2}
	third1 := TreeNode{Val: 3}
	first1.Left = &second1
	first1.Right = &third1

	fmt.Println(checkSubTree(&first, &first1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkSubTree(t1 *TreeNode, t2 *TreeNode) bool {
	sStr := preOrder(t1)
	tStr := preOrder(t2)
	return strings.Contains(sStr, tStr)
}

func preOrder(root *TreeNode) string {
	if root == nil {
		return ""
	}
	res := "!"
	stack := make([]*TreeNode, 0)
	temp := root
	for {
		for temp != nil {
			res += strconv.Itoa(temp.Val)
			res += "!"
			stack = append(stack, temp)
			temp = temp.Left
		}
		res += "#!"
		if len(stack) > 0 {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			temp = node.Right
		} else {
			break
		}
	}
	return res
}
