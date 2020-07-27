package main

import "fmt"

func main() {
	first := TreeNode{Val: 1}
	second := TreeNode{Val: 2}
	third := TreeNode{Val: 3}
	first.Left = &second
	first.Right = &third
	fmt.Println(levelOrder(&first))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 剑指offer_面试题32-III_从上到下打印二叉树III
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	temp := make([][]int, 0)
	dfs(root, &temp, 0)
	return temp
}

func dfs(root *TreeNode, temp *[][]int, level int) {
	if root == nil {
		return
	}
	if len(*temp)-1 < level {
		*temp = append(*temp, make([]int, 0))
	}
	if level%2 == 0 {
		(*temp)[level] = append((*temp)[level], root.Val)
	} else {
		(*temp)[level] = append([]int{root.Val}, (*temp)[level]...)
	}
	level = level + 1
	dfs(root.Left, temp, level)
	dfs(root.Right, temp, level)
}
