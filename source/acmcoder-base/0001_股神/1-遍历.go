package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	res := getResult(n)
	fmt.Println(res)
}

func getResult(n int) int {
	if n <= 2 {
		return n
	}
	count := 0 // 统计下跌多少次
	target := n
	// n-2-3-4-...-
	// 每减去1次，碰到的下跌次数+1
	for i := 2; i <= n; i++ {
		if target-i > 0 {
			target = target - i
			count++
		} else {
			break
		}
	}
	return n - 2*count
}
