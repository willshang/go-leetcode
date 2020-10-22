package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(translateNum(10022))
	fmt.Println(string(48))
}

// 剑指offer_面试题46_把数字翻译成字符串
// f(i)=f(i-2)+f(i-1)
// f(i)=f(i-1)
func translateNum(num int) int {
	if num < 0 {
		return 0
	}
	str := strconv.Itoa(num)
	arr := make([]int, len(str))
	for i := 0; i < len(str); i++ {
		arr[i] = int(str[i] - '0')
	}
	dp := make([]int, len(str)+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i < len(str)+1; i++ {
		if arr[i-2] != 0 && (arr[i-2]*10+arr[i-1] <= 25) {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(str)]
}
