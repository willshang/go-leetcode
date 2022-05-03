package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode_lcp52_二叉搜索树染色
var arr []int

func getNumber(root *TreeNode, ops [][]int) int {
	res := 0 // 红色数量
	arr = make([]int, 0)
	dfs(root)
	for i := 0; i < len(arr); i++ {
		for j := len(ops) - 1; j >= 0; j-- {
			t, left, right := ops[j][0], ops[j][1], ops[j][2]
			if left <= arr[i] && arr[i] <= right {
				res = res + t // 红色+1，蓝色+0
				break
			}
		}
	}
	return res // 红色数量
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	arr = append(arr, root.Val)
	dfs(root.Right)
}
