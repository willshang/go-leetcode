package main

func main() {

}

// leetcode926_将字符串翻转到单调递增
func minFlipsMonoIncr(S string) int {
	n := len(S)
	dpA := make([]int, n) // 0 结尾
	dpB := make([]int, n) // 1 结尾
	if S[0] == '1' {
		dpA[0] = 1
	} else {
		dpB[0] = 1
	}
	for i := 1; i < n; i++ {
		if S[i] == '1' {
			dpA[i] = dpA[i-1] + 1            // 需要改为0
			dpB[i] = min(dpB[i-1], dpA[i-1]) // 1结尾和0结尾的最小值
		} else {
			dpA[i] = dpA[i-1]                    // 不需要改为0
			dpB[i] = min(dpB[i-1], dpA[i-1]) + 1 // 1结尾和0结尾的最小值+1
		}
	}
	return min(dpA[n-1], dpB[n-1])
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
