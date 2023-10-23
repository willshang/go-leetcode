package main

import "fmt"

func main() {
	fmt.Println(numberOfSets(4, 2))
	fmt.Println(numberOfSets(30, 7))
}

// 1621
// 参考：费马小定理
// https://leetcode-cn.com/problems/number-of-sets-of-k-non-overlapping-line-segments/solution/dong-tai-gui-hua-shu-xue-fa-si-lu-fen-xi-sagz/
var mod = 1000000007

func numberOfSets(n int, k int) int {
	// 共享k-1个点+n个点 => n+k-1个点
	// 共 n+k−1 个数中选择 2k 个
	return C(n+k-1, 2*k)
}

func C(n, m int) int {
	a := multiMod(n, n-m+1)
	b := multiMod(m, 1)
	return a * powMod(b, mod-2) % mod
}

func multiMod(n, m int) int {
	res := 1
	for i := m; i <= n; i++ {
		res = (res * i) % mod
	}
	return res
}

func powMod(a, b int) int {
	res := 1
	for b > 0 {
		if b%2 == 1 {
			res = (res * a) % mod
		}
		a = a * a % mod
		b = b / 2
	}
	return res
}
