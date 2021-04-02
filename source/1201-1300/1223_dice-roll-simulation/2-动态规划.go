package main

func main() {

}

func dieSimulator(n int, rollMax []int) int {
	dp := [7][16]int{} // 以j结尾,出现k次
	for j := 1; j <= 6; j++ {
		dp[j][1] = 1
	}
	for i := 2; i <= n; i++ {
		temp := [7][16]int{}
		for j := 1; j <= 6; j++ {
			for k := 1; k <= rollMax[j-1] && k <= i; k++ {
				if k == 1 { // 不同情况
					for a := 1; a <= 6; a++ {
						if a != j { // 考虑不同的情况
							for b := 1; b <= rollMax[a-1] && b <= i-1; b++ {
								temp[j][k] = (temp[j][k] + dp[a][b]) % 1000000007
							}
						}
					}
				} else { // 直接同上一个
					temp[j][k] = dp[j][k-1]
				}
			}
		}
		dp = temp
	}
	res := 0
	for j := 1; j <= 6; j++ {
		for k := 1; k <= 15; k++ {
			res = (res + dp[j][k]) % 1000000007
		}
	}
	return res
}
