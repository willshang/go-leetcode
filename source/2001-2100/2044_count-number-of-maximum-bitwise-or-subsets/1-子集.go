package main

func main() {

}

func countMaxOrSubsets(nums []int) int {
	n := len(nums)
	total := 1 << n
	sum := make([]int, total)
	for i := 0; i < n; i++ { // 每次添加1个
		count := 1 << i
		for j := 0; j < count; j++ {
			sum[count|j] = sum[j] | nums[i] // 按位或运算：j前面补1=>子集和加上tasks[i]
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
