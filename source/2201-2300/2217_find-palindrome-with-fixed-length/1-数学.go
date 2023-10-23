package main

import (
	"math"
	"strconv"
)

func main() {

}

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
	arr := []byte(strconv.Itoa(num))
	if intLength%2 == 0 { // 偶数
		for i := len(arr) - 1; i >= 0; i-- {
			arr = append(arr, arr[i])
		}
	} else {
		for i := len(arr) - 2; i >= 0; i-- {
			arr = append(arr, arr[i])
		}
	}
	res, _ := strconv.ParseInt(string(arr), 10, 64)
	return res
}
