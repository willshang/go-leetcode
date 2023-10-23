package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int

func averageOfSubtree(root *TreeNode) int {
	res = 0
	dfs(root)
	return res
}

func dfs(root *TreeNode) (sum, count int) {
	sum, count = root.Val, 1
	if root.Left != nil {
		sL, cL := dfs(root.Left)
		sum = sum + sL
		count = count + cL
	}
	if root.Right != nil {
		sR, cR := dfs(root.Right)
		sum = sum + sR
		count = count + cR
	}
	if sum/count == root.Val {
		res++
	}
	return sum, count
}
