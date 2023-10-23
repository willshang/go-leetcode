package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(largestPalindrome(2))
}

// leetcode479_最大回文数乘积
func largestPalindrome(n int) int {
	if n == 1 {
		return 9
	}
	last := int(math.Pow10(n)) - 1 // 10^n-1 ：n=3 999
	first := last / 10             // 10^(n-1)-1 n=3 99
	for i := last; i > first; i-- {
		target := makePalindrome(i) // 依次生成对应的递减回文数：如98=>9889
		for j := last; j > first && target < j*j; j-- {
			if target%j == 0 { // 该回文数满足要求
				return target % 1337
			}
		}
	}
	return 0
}

func makePalindrome(num int) int {
	str := strconv.Itoa(num)
	arr := []byte(str)
	for i := len(str) - 1; i > -1; i-- {
		arr = append(arr, str[i])
	}
	res, _ := strconv.Atoi(string(arr))
	return res
}

/*var res = []int{9, 987, 123, 597, 677, 1218, 877, 475}
func largestPalindrome(n int) int {
	return res[n-1]
}*/
