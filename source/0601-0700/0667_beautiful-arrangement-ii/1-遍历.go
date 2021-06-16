package main

import "fmt"

func main() {
	fmt.Println(constructArray(10, 3))
	fmt.Println(constructArray(3, 1))
}

// leetcode667_优美的排列II
func constructArray(n int, k int) []int {
	if n == k {
		return nil
	}
	res := make([]int, n)
	// 构建等差数列为1：共n-k个数=>1
	for i := 1; i <= n-k; i++ {
		res[i-1] = i
	}
	// 构建交错队列：最大值和最小值交错出现，这样差值各不相同
	// n=10, k=7 => [1 2 3 4 5 6 7 10 8 9]
	// 剩下k个数（与等差数列相连）：共k个差值，依次1、2、3、...，去除1后共k-1个差值
	left := n - k + 1
	right := n
	count := 0
	for i := n - k + 1; i <= n; i++ {
		if count%2 == 1 {
			res[i-1] = left
			left++
		} else {
			res[i-1] = right
			right--
		}
		count++
	}
	return res
}
