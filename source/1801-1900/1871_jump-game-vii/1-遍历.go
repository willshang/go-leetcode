package main

func main() {

}

// leetcode1871_跳跃游戏VII
func canReach(s string, minJump int, maxJump int) bool {
	n := len(s)
	if s[n-1] == '1' {
		return false
	}
	minDis, maxDis := 0, 0 // 更新最后可达到的最小+最大坐标的范围
	for i := 0; i < n-1; i++ {
		if s[i] == '0' && minDis <= i && i <= maxDis {
			minDis = i + minJump
			maxDis = i + maxJump
		}
		if minDis <= n-1 && n-1 <= maxDis {
			return true
		}
	}
	return false
}
