package main

func main() {

}

// leetcode2212_射箭比赛中的最大得分
func maximumBobPoints(numArrows int, aliceArrows []int) []int {
	n := len(aliceArrows)
	total := 1 << n // 所有状态
	maxValue := 0
	maxState := 0
	for state := 0; state < total; state++ { // 枚举所有状态
		sum := 0
		score := 0
		for i := 0; i < n; i++ {
			if (state>>i)&1 > 0 {
				sum = sum + aliceArrows[i] + 1 // 箭的数量
				score = score + i              // 得分
			}
		}
		if sum <= numArrows && score > maxValue {
			maxValue = score
			maxState = state
		}
	}
	res := make([]int, n)
	for i := 0; i < n; i++ {
		if (maxState>>i)&1 > 0 {
			res[i] = aliceArrows[i] + 1
			numArrows = numArrows - res[i]
		}
	}
	res[0] = res[0] + numArrows // 剩下放在任何1个里面即可
	return res
}
