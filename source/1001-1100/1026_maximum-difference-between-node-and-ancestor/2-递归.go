package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int

func maxAncestorDiff(root *TreeNode) int {
	res = 0
	if root == nil {
		return 0
	}
	dfs(root, []int{})
	return res
}

func dfs(root *TreeNode, arr []int) {
	if root == nil {
		return
	}
	for i := 0; i < len(arr); i++ {
		if abs(arr[i], root.Val) > res {
			res = abs(arr[i], root.Val)
		}
	}
	arr = append(arr, root.Val)
	dfs(root.Left, arr)
	dfs(root.Right, arr)
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
