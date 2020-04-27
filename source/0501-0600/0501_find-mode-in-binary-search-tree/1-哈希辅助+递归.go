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

// leetcode501_二叉搜索树中的众数
func findMode(root *TreeNode) []int {
	m := map[int]int{}
	dfs(root, m)
	max := -1
	res := make([]int, 0)
	for i, v := range m {
		if max <= v {
			if max < v {
				max = v
				res = res[0:0]
			}
			res = append(res, i)
		}
	}
	return res
}

func dfs(root *TreeNode, rec map[int]int) {
	if root == nil {
		return
	}
	rec[root.Val]++
	dfs(root.Left, rec)
	dfs(root.Right, rec)
}
