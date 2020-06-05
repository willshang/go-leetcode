package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(twoSum(3))
}

// 剑指offer_面试题60_n个骰子的点数
// f(n,k)=f(n−1,k−1)+f(n−1,k−2)+f(n−1,k−3)+f(n−1,k−4)+f(n−1,k−5)+f(n−1,k−6)
func twoSum(n int) []float64 {
	res := make([]float64, 0)
	if n < 1 {
		return res
	}
	arr := make([]int, 6*n+1)
	for i := 1; i <= 6; i++ {
		arr[i] = 1
	}
	for i := 2; i <= n; i++ {
		for j := 6 * i; j >= i; j-- {
			arr[j] = 0
			for k := 1; k <= 6; k++ {
				if j-k >= i-1 {
					arr[j] = arr[j] + arr[j-k]
				}
			}
		}
	}
	total := math.Pow(float64(6), float64(n))
	for i := n; i <= 6*n; i++ {
		ratio := float64(arr[i]) / total
		res = append(res, ratio)
	}
	return res
}
