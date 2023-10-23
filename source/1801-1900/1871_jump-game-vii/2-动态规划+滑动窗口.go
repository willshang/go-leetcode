package main

func main() {

}

func canReach(s string, minJump int, maxJump int) bool {
	n := len(s)
	if s[n-1] == '1' {
		return false
	}
	dp := make([]bool, n) // dp[i] => 下标i是否可达
	dp[0] = true
	count := 1 // 滑动窗口里面可达数量
	for i := minJump; i < n; i++ {
		if s[i] == '0' && count > 0 { // 当前可达
			dp[i] = true
		}
		if maxJump <= i && dp[i-maxJump] == true { // 滑动窗口左端点:-1
			count--
		}
		if dp[i-minJump+1] == true { // 滑动窗口右端点：+1
			count++
		}
	}
	return dp[n-1]
}
