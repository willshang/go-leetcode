package main

func main() {

}

func canDistribute(nums []int, quantity []int) bool {
	n := len(quantity)
	total := 1 << n
	sum := make([]int, total)
	for i := 0; i < n; i++ {
		count := 1 << i
		for j := 0; j < count; j++ {
			sum[count|j] = sum[j] + quantity[i] // 按位或运算：j前面补1=>子集和加上tasks[i]
		}
	}
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	dp := make([][]bool, len(m)+1) // dp[i][j] => m的前i个元素是否满足集合状态为j的顾客组合
	for i := 0; i <= len(m); i++ {
		dp[i] = make([]bool, total)
		dp[i][0] = true
	}
	i := 0 // experience map 是随机的
	for _, v := range m {
		for i := 0; i < len(dp[v]); i++ {
			if dp[v][i] == true {
			}
		}
	}
	return dp[len(m)][total-1]
}
