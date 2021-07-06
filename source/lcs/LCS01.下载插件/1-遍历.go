package main

import "fmt"

func main() {
	fmt.Println(leastMinutes(8))
}

// leetcode_lcs01_下载插件
func leastMinutes(n int) int {
	res := 0
	for i := 1; i < n; i = i * 2 {
		res++
	}
	return res + 1
}
