package main

import "fmt"

func main() {
	fmt.Println(numWays("10101"))
	fmt.Println(numWays("1001"))
	fmt.Println(numWays("0000"))
	fmt.Println(numWays("100100010100110"))
}

// leetcode1573_分割字符串的方案数
func numWays(s string) int {
	arr := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			arr = append(arr, i)
		}
	}
	if len(arr)%3 != 0 {
		return 0
	}
	if len(arr) == 0 {
		return ((len(s) - 2) * (len(s) - 1) / 2) % 1000000007
	}
	left, right := arr[len(arr)/3]-arr[len(arr)/3-1], arr[len(arr)/3*2]-arr[len(arr)/3*2-1]
	return left * right % 1000000007
}
