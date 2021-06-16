package main

func main() {

}

var mod = 1000000007

func numOfWays(n int) int {
	arr := make([]int, 0) // 保存所有满足条件的排列
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				if i != j && j != k {
					arr = append(arr, i*100+j*10+k)
				}
			}
		}
	}
	length := len(arr)
	judgeArr := make([][]int, length) // 相邻关系判断：1代表相邻行可以相邻
	for i := 0; i < length; i++ {
		judgeArr[i] = make([]int, length)
		for j := 0; j < length; j++ {
			if arr[i]/100 != arr[j]/100 &&
				arr[i]/10%10 != arr[j]/10%10 &&
				arr[i]%10 != arr[j]%10 {
				judgeArr[i][j] = 1
			}
		}
	}
	// 上面是预处理
	dp := make([][]int, n+1) // dp[i][j]表示：i个x3网格，最后一行是呈现的是j的 方案数
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, length)
	}
	for j := 0; j < length; j++ {
		dp[1][j] = 1 // 第一行可以使用任何类型
	}
	for i := 2; i <= n; i++ {
		for j := 0; j < length; j++ {
			for k := 0; k < length; k++ {
				if judgeArr[j][k] == 1 {
					dp[i][j] = (dp[i][j] + dp[i-1][k]) % mod
				}
			}
		}
	}
	res := 0
	for j := 0; j < length; j++ {
		res = (res + dp[n][j]) % mod
	}
	return res
}
