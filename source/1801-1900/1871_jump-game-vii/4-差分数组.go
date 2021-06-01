package main

func main() {

}

func canReach(s string, minJump int, maxJump int) bool {
	n := len(s)
	if s[n-1] == '1' {
		return false
	}
	arr := make([]int, n+1)
	arr[0] = 1
	arr[1] = -1
	sum := 0
	for i := 0; i < n; i++ {
		sum = sum + arr[i]
		if s[i] == '0' && sum > 0 {
			arr[min(i+minJump, n)]++
			arr[min(i+maxJump+1, n)]--
		}
	}
	return sum > 0 // 计算范围内可达到最后有几个坐标满足条件
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
