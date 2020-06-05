package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(twoSum(3))
}

func twoSum(n int) []float64 {
	res := make([]float64, 0)
	if n < 1 {
		return res
	}
	arr := [2][]int{}
	arr[0] = make([]int, 6*n+1)
	arr[1] = make([]int, 6*n+1)
	flag := 0
	for i := 1; i <= 6; i++ {
		arr[flag][i] = 1
	}
	for k := 2; k <= n; k++ {
		for i := 0; i < k; i++ {
			arr[1-flag][i] = 0
		}
		for i := k; i <= 6*n; i++ {
			arr[1-flag][i] = 0
			// 当前轮的第N位等于前一个数组第N-1,N-2,N-3,N-4,N-5,N-6位之和
			for j := 1; j <= i && j <= 6; j++ {
				arr[1-flag][i] += arr[flag][i-j]
			}
		}
		flag = 1 - flag
	}
	total := math.Pow(float64(6), float64(n))
	for i := n; i <= 6*n; i++ {
		ratio := float64(arr[flag][i]) / total
		res = append(res, ratio)
	}
	return res
}
