package main

import (
	"fmt"
)

func main() {
	fmt.Println(combine(4, 4))
}

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	arr := make([]int, 0)
	for i := 1; i <= k; i++ {
		arr = append(arr, 0)
	}
	i := 0
	for i >= 0 {
		arr[i]++
		if arr[i] > n {
			i--
		} else if i == k-1 {
			temp := make([]int, k)
			copy(temp, arr)
			res = append(res, temp)
		} else {
			i++
			arr[i] = arr[i-1]
		}
	}
	return res
}
