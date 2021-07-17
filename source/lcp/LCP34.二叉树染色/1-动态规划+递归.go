package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode_lcp34_二叉树染色
func maxValue(root *TreeNode, k int) int {
	dp := dfs(root, k)
	return maxArr(dp)
}

func dfs(root *TreeNode, k int) []int {
	dp := make([]int, k+1) // dp[i]表示，染色数为i的最大值
	if root == nil {
		return dp
	}
	left := dfs(root.Left, k)
	right := dfs(root.Right, k)
	dp[0] = maxArr(left) + maxArr(right) // 当前节点不染色
	for i := 1; i <= k; i++ {            // 当前节点染色
		for j := 0; j < i; j++ {
			dp[i] = max(dp[i], left[j]+right[i-1-j]+root.Val)
		}
	}
	return dp
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxArr(arr []int) int {
	res := 0
	for i := 0; i < len(arr); i++ {
		res = max(res, arr[i])
	}
	return res
}
