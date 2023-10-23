package main

func main() {

}

func canReach(s string, minJump int, maxJump int) bool {
	n := len(s)
	if s[n-1] == '1' {
		return false
	}
	dp := make([]int, n) // dp[i] => 下标i是否可达
	dp[0] = 1
	arr := make([]int, n)
	for i := 0; i < minJump; i++ { // 前缀和从minJump开始
		arr[i] = 1
	}
	for i := minJump; i < n; i++ {
		left := i - maxJump
		right := i - minJump
		if s[i] == '0' {
			var total int
			if left <= 0 {
				total = arr[right]
			} else {
				total = arr[right] - arr[left-1]
			}
			if total > 0 { // 通过前缀和计算，范围内存在多少满足条件的坐标
				dp[i] = 1
			}
		}
		arr[i] = arr[i-1] + dp[i]
	}
	return dp[n-1] > 0
}
