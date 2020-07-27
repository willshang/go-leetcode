package main

import "fmt"

func main() {
	fmt.Println(getRow(3))
}

// leetcode 119 杨辉三角 II
func getRow(rowIndex int) []int {
	res := make([]int, rowIndex+1)
	res[0] = 1
	if rowIndex == 0 {
		return res
	}

	// 公式
	// C(n,k）= n! /(k! * (n-k)!)
	// C(n,k) = (n-k+1)/k * C(n,k-1)
	for i := 1; i <= rowIndex; i++ {
		res[i] = res[i-1] * (rowIndex - i + 1) / i
	}
	return res
}
