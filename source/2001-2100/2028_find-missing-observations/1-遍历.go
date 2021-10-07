package main

func main() {

}

// leetcode2028_找出缺失的观测数据
func missingRolls(rolls []int, mean int, n int) []int {
	sum := 0
	m := len(rolls)
	for i := 0; i < m; i++ {
		sum = sum + rolls[i]
	}
	total := mean * (n + m)
	left := total - sum
	if left > n*6 || left < n {
		return nil
	}
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = left / n // 平均分配
	}
	for i := 0; i < left%n; i++ {
		res[i] = res[i] + 1
	}
	return res
}
