package main

func main() {

}

func countEval(s string, result int) int {
	n := len(s)
	// dp[i][j][0/1] => s[i:j+1]结果为0/1的方法数
	dp := make([][][2]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][2]int, n)
		if i%2 == 0 {
			dp[i][i][int(s[i]-'0')]++
		}
	}
	for length := 2; length <= n; length = length + 2 { // 枚举长度
		for i := 0; i <= n-length; i = i + 2 { // 枚举起点
			j := i + length                    // 确定终点
			for k := i + 1; k < j; k = k + 2 { // 枚举操作符
				for a := 0; a <= 1; a++ {
					for b := 0; b <= 1; b++ {
						temp := getValue(a, b, s[k])
						dp[i][j][temp] = dp[i][j][temp] +
							dp[i][k-1][a]*dp[k+1][j][b]
					}
				}
			}
		}
	}
	return dp[0][n-1][result]
}

func getValue(a, b int, op byte) int {
	if op == '&' {
		return a & b
	} else if op == '|' {
		return a | b
	}
	return a ^ b
}
