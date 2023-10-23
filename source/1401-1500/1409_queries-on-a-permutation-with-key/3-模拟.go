package main

import "fmt"

func main() {
	fmt.Println(processQueries([]int{3, 1, 2, 1}, 5))
}

func processQueries(queries []int, m int) []int {
	n := len(queries)
	res := make([]int, n)
	arr := make([]int, m+1) // 存放下标
	for i := 1; i <= m; i++ {
		arr[i] = i - 1
	}
	for i := 0; i < n; i++ {
		x := queries[i]
		index := arr[x]
		res[i] = index
		for j := 1; j <= m; j++ {
			if arr[j] < index { // 下标小于目标值
				arr[j]++ // 前面的后移
			}
		}
		arr[x] = 0 // 移到第一位
	}
	return res
}
