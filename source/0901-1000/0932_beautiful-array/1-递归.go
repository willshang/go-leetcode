package main

import "fmt"

func main() {
	fmt.Println(beautifulArray(10))
}

var m map[int][]int

func beautifulArray(N int) []int {
	m = make(map[int][]int)
	return dfs(N)
}

func dfs(n int) []int {
	if v, ok := m[n]; ok {
		return v
	}
	res := make([]int, n)
	if n == 1 {
		res[0] = 1
	} else {
		// A[k] * 2 = A[i] + A[j], i < k < j
		// 要想等式恒不成立，一个简单的办法就是让i部分的数都是奇数，j部分的数都是偶数。
		index := 0
		for _, v := range dfs((n + 1) / 2) { // 1-N中有(N+1)/2个奇数
			res[index] = 2*v - 1
			index++
		}
		for _, v := range dfs(n / 2) { // 1-N中有N/2个偶数
			res[index] = 2 * v
			index++
		}
	}
	m[n] = res
	return res
}
