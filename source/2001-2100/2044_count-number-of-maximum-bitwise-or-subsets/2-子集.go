package main

func main() {

}

func countMaxOrSubsets(nums []int) int {
	n := len(nums)
	total := 1 << n
	sum := make([]int, total)
	for i := 0; i < total; i++ { // 枚举状态
		for j := 0; j < n; j++ { // 枚举该位
			if (i & (1 << j)) > 0 {
				sum[i] = sum[i] | nums[j]
			}
		}
	}
	maxValue := 0
	for i := 0; i < total; i++ {
		maxValue = max(maxValue, sum[i])
	}
	res := 0
	for i := 0; i < total; i++ {
		if sum[i] == maxValue {
			res++
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
