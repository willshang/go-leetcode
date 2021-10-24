package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	n := len(arr)
	total := 1 << n
	sum := make([]int, total)
	for i := 0; i < total; i++ { // 枚举状态
		for j := 0; j < n; j++ { // 枚举该位
			if (i & (1 << j)) > 0 {
				sum[i] = sum[i] + arr[j]
			}
		}
	}
	for i := 0; i < len(sum); i++ {
		fmt.Printf("%05b %d\n", i, sum[i])
	}
}
