package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	n := len(arr)
	total := 1 << n
	sum := make([]int, total)
	for i := 0; i < n; i++ {
		count := 1 << i
		for j := 0; j < count; j++ {
			sum[count|j] = sum[j] + arr[i] // 按位或运算：j前面补1=>子集和加上arr[i]
		}
	}
	for i := 0; i < len(sum); i++ {
		fmt.Printf("%05b %d\n", i, sum[i])
	}
}
