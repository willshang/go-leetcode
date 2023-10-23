package main

import "fmt"

func main() {
	fmt.Println(preimageSizeFZF(0))
}

// leetcode793_阶乘函数后K个零
func preimageSizeFZF(k int) int {
	left, right := 0, 5*k+1
	for left < right {
		mid := left + (right-left)/2
		target := trailingZeroes(mid)
		if target == k {
			return 5 // 能找到一定会存在连续5个数的阶乘末尾0的个数为k
		} else if target < k {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return 0 // 不存在返回0
}

// leetcode 172.阶乘后的零
// n!末尾0的个数
func trailingZeroes(n int) int {
	result := 0
	for n >= 5 {
		n = n / 5
		result = result + n
	}
	return result
}
