package main

import "fmt"

func main() {
	fmt.Println(candy([]int{1, 0, 2}))
}

func candy(ratings []int) int {
	arr := make([]int, len(ratings))
	// 规则1：每个孩子至少分配到 1 个糖果。
	for i := 0; i < len(arr); i++ {
		arr[i] = 1
	}
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] && arr[i] <= arr[i-1] {
			arr[i] = arr[i-1] + 1
		}
	}
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && arr[i] <= arr[i+1] {
			arr[i] = arr[i+1] + 1
		}
	}
	res := 0
	for i := 0; i < len(arr); i++ {
		res = res + arr[i]
	}
	return res
}
