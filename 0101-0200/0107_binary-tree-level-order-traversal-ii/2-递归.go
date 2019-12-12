package main

import "fmt"

func main() {
	root := TreeNode{}
	root.Val = 1

	left := TreeNode{}
	left.Val = 2

	right := TreeNode{}
	right.Val = 3

	root.Left = &left
	root.Right = &right
	res := levelOrderBottom(&root)

	for _, v := range res {
		fmt.Println(v)
	}

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode 107 二叉树的层次遍历 II
func levelOrderBottom(root *TreeNode) [][]int {
	result := make([][]int, 0)
	level := 0
	if root == nil {
		return result
	}

	orderBottom(root, &result, level)

	left, right := 0, len(result)-1
	for left < right {
		result[left], result[right] = result[right], result[left]
		left++
		right--
	}
	return result
}

func orderBottom(root *TreeNode, result *[][]int, level int) {
	if root == nil {
		return
	}
	if len(*result) > level {
		fmt.Println(level, result, root.Val)
		(*result)[level] = append((*result)[level], root.Val)
	} else {
		*result = append(*result, []int{root.Val})
	}
	orderBottom(root.Left, result, level+1)
	orderBottom(root.Right, result, level+1)
}
