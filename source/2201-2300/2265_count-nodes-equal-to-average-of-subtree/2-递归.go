package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode2265_统计值等于子树平均值的节点数
var res int

func averageOfSubtree(root *TreeNode) int {
	res = 0
	dfs(root)
	return res
}

func dfs(root *TreeNode) (sum, count int) {
	if root == nil {
		return 0, 0
	}
	sL, cL := dfs(root.Left)
	sR, cR := dfs(root.Right)
	sum = root.Val + sL + sR
	count = 1 + cL + cR
	if sum/count == root.Val {
		res++
	}
	return sum, count
}
