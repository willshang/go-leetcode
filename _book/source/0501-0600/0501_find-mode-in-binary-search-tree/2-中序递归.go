package main

import "fmt"

func main() {
	first := TreeNode{}
	first.Val = 19
	second := TreeNode{}
	second.Val = 5
	third := TreeNode{}
	third.Val = 19
	fourth := TreeNode{}
	fourth.Val = 5

	first.Right = &second
	second.Left = &third
	second.Right = &fourth

	fmt.Println(findMode(&first))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var max int
var res []int
var cur int
var count int

func findMode(root *TreeNode) []int {
	res = make([]int, 0)
	max, cur, count = 0, 0, 0
	dfs(root)
	return res
}

// 中序遍历保证利用二叉搜索树的性质，得到的结果是升序的
func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	if root.Val != cur {
		count = 0
	}
	count++
	if max < count {
		max = count
		res = []int{root.Val}
	} else if max == count {
		res = append(res, root.Val)
	}
	cur = root.Val
	dfs(root.Right)
}
