package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sum int
var target int

func maxProduct(root *TreeNode) int {
	sum = 0
	target = 0
	dfsSum(root)
	dfs(root)
	return target * (sum - target) % 1000000007
}

func dfsSum(root *TreeNode) {
	if root == nil {
		return
	}
	sum = sum + root.Val
	dfsSum(root.Left)
	dfsSum(root.Right)
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left)
	right := dfs(root.Right)
	total := left + right + root.Val
	if abs(sum-2*total) < abs(sum-2*target) {
		target = total
	}
	return total
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
