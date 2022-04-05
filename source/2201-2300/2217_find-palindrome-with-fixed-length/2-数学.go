package main

import (
	"math"
)

func main() {

}

// leetcode2217_找到指定长度的回文数
func kthPalindrome(queries []int, intLength int) []int64 {
	n := (intLength + 1) / 2          // 回文前半部分的长度：12321=>长度5=>前半部分长度3
	start := int(math.Pow10(n-1) - 1) // n长度对应的回文数量下限 ：2=>10^2-1=99
	limit := int(math.Pow10(n) - 1)   // n长度对应的回文数量上限
	res := make([]int64, 0)
	for i := 0; i < len(queries); i++ {
		if start+queries[i] > limit {
			res = append(res, -1)
			continue
		}
		res = append(res, getKthPalindrome(intLength, start+queries[i]))
	}
	return res
}

func getKthPalindrome(intLength, num int) int64 {
	temp := num
	if intLength%2 == 1 {
		temp = temp / 10
	}
	res := int64(num)
	for ; temp > 0; temp = temp / 10 {
		res = res*10 + int64(temp%10)
	}
	return res
}
