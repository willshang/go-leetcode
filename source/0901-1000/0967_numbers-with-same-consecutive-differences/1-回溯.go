package main

import "fmt"

func main() {
	fmt.Println(numsSameConsecDiff(3, 7))
	fmt.Println(numsSameConsecDiff(2, 0))
}

// leetcode967_连续差相同的数字
var res []int

func numsSameConsecDiff(n int, k int) []int {
	res = make([]int, 0)
	for i := 1; i <= 9; i++ {
		dfs(i, n, k, i)
	}
	return res
}

func dfs(index, n, k, value int) {
	if n == 1 {
		res = append(res, value)
		return
	}
	a := index + k
	if 0 <= a && a <= 9 && k != 0 {
		value = value*10 + a
		dfs(a, n-1, k, value)
		value = (value - a) / 10
	}
	b := index - k
	if 0 <= b && b <= 9 {
		value = value*10 + b
		dfs(b, n-1, k, value)
		value = (value - b) / 10
	}
}
