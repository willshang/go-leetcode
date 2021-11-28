package main

func main() {

}

// leetcode1483_树节点的第K个祖先
type TreeAncestor struct {
	dp [][]int
}

func Constructor(n int, parent []int) TreeAncestor {
	m := 0
	for i := 1; i <= n; i = i * 2 {
		m++
	}
	dp := make([][]int, n) // dp[i][j] => i的第2^j个父节点
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = -1
		}
	}
	for i := 0; i < n; i++ {
		dp[i][0] = parent[i]
	}
	for i := 0; i < n; i++ {
		for j := 1; j < m; j++ {
			if dp[i][j-1] == -1 { // 没有父节点
				continue
			}
			prev := dp[i][j-1] // 状态转移方程：dp[i][j] = dp[prev][j-1]
			dp[i][j] = dp[prev][j-1]
		}
	}
	return TreeAncestor{dp: dp}
}

func (this *TreeAncestor) GetKthAncestor(node int, k int) int {
	for i := 0; i < 16; i++ {
		if k&(1<<i) > 0 {
			node = this.dp[node][i]
		}
		if node == -1 {
			break
		}
	}
	return node
}
