package main

func main() {

}

// leetcode390_消除游戏
func lastRemaining(n int) int {
	return dfs(n, 1)
}

func dfs(n int, isLeftToRight int) int {
	if n == 1 {
		return 1
	}
	if n%2 == 1 { // 奇数
		return 2 * dfs(n/2, 1-isLeftToRight)
	}
	if isLeftToRight == 1 { // 偶数从左往右
		return 2 * dfs(n/2, 1-isLeftToRight)
	}
	// 偶数从右往左，如1，2，3，4，5，6，7，8 => 1，3，5，7
	return 2*dfs(n/2, 1-isLeftToRight) - 1
}
