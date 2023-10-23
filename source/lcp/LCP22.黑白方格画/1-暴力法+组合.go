package main

import "fmt"

func main() {
	fmt.Println(paintingPlan(2, 2))
}

// leetcode_lcp22_黑白方格画
func paintingPlan(n int, k int) int {
	if k == n*n || k == 0 { // 全部涂满或者不涂只有1种方案
		return 1
	}
	if k < n { // 最少大于等于n
		return 0
	}
	res := 0
	// 暴力枚举行和列
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			a := i * n
			b := j * n
			if a+b-i*j == k {
				res = res + C(n, i)*C(n, j) // 求组合数
			}
		}
	}
	return res
}

func C(n, m int) int {
	a := 1
	for i := 1; i <= m; i++ {
		a = a * (n - i + 1)
	}

	b := 1
	for i := 1; i <= m; i++ {
		b = b * i
	}
	return a / b
}
