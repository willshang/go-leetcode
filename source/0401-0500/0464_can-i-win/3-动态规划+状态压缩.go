package main

func main() {

}

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	a, b := maxChoosableInteger, desiredTotal
	if a*(a+1)/2 < b {
		return false
	}
	total := 1 << a
	dp := make([]bool, total)
	for i := total - 1; i >= 0; i-- { // 枚举状态
		sum := 0
		for k := 0; k < a; k++ { // 状态和
			if i&(1<<k) > 0 { // i：对应状态为1位置上和
				sum = sum + (k + 1)
			}
		}
		for k := 0; k < a; k++ {
			if i&(1<<k) > 0 {
				continue
			}
			prev := i | (1 << k)
			// >=剩下值，或者之前为false
			if k+1 >= desiredTotal-sum || dp[prev] == false {
				dp[i] = true
			}
		}
	}
	return dp[0]
}
