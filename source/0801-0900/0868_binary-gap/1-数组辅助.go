package main

import "fmt"

func main() {
	fmt.Println(binaryGap(22))
}

func binaryGap(N int) int {
	arr := make([]int, 0)
	index := 0
	for N > 0 {
		if N%2 == 1 {
			arr = append(arr, index)
		}
		index++
		N = N / 2
	}
	res := 0
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] > res {
			res = arr[i+1] - arr[i]
		}
	}
	return res
}
