package main

import "fmt"

func main() {
	fmt.Println(getXORSum([]int{12}, []int{4}))
}

// leetcode1835_所有数对按位与结果的异或和
func getXORSum(arr1 []int, arr2 []int) int {
	a, b := 0, 0
	for i := 0; i < len(arr1); i++ {
		a = a ^ arr1[i]
	}
	for i := 0; i < len(arr2); i++ {
		b = b ^ arr2[i]
	}
	return a & b // 位与操作：相同位数都是1，则该结果返回1
}
