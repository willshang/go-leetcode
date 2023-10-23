package main

import "fmt"

func main() {
	fmt.Println(platesBetweenCandles("**|**|***|", [][]int{
		{2, 5},
		{5, 9},
	}))

	fmt.Println(platesBetweenCandles("***|**|*****|**||**|*", [][]int{
		{1, 17},
		{4, 5},
	}))
}

// leetcode2055_蜡烛之间的盘子
func platesBetweenCandles(s string, queries [][]int) []int {
	n := len(s)
	res := make([]int, len(queries))
	sum := make([]int, n+1)
	left, right := make([]int, n), make([]int, n)
	prev := -1
	for i := 0; i < n; i++ {
		sum[i+1] = sum[i]
		if s[i] == '|' { // 蜡烛
			prev = i
		} else {
			sum[i+1]++
		}
		left[i] = prev
	}
	prev = n
	for i := n - 1; i >= 0; i-- {
		if s[i] == '|' { // 蜡烛
			prev = i
		}
		right[i] = prev
	}
	for i := 0; i < len(queries); i++ {
		a, b := right[queries[i][0]], left[queries[i][1]] // 位子：右边第一个蜡烛，左边第一个蜡烛
		if a < b {                                        // 满足的条件下
			res[i] = sum[b] - sum[a]
		}
	}
	return res
}
