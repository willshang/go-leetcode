package main

import "fmt"

func main() {
	fmt.Println(circularPermutation(2, 3))
}

func circularPermutation(n int, start int) []int {
	total := 1 << n
	res := []int{0, 1}
	for i := 2; i <= n; i++ {
		// 计算格雷码：leetcode89.格雷编码
		for j := len(res) - 1; j >= 0; j-- {
			res = append(res, res[j]+(1<<(i-1)))
		}
	}
	index := 0
	for i := 0; i < total; i++ {
		if res[i] == start {
			index = i
			break
		}
	}
	return append(res[index:], res[:index]...)
}
