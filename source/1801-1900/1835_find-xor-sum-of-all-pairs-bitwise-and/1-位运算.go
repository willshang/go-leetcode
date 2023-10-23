package main

import "fmt"

func main() {
	fmt.Println(getXORSum([]int{12}, []int{4}))
}

func getXORSum(arr1 []int, arr2 []int) int {
	res := 0
	n, m := len(arr1), len(arr2)
	for i := 31; i >= 0; i-- {
		a, b := 0, 0
		for j := 0; j < n; j++ {
			if (arr1[j] & (1 << i)) > 0 {
				a++
			}
		}
		for j := 0; j < m; j++ {
			if (arr2[j] & (1 << i)) > 0 {
				b++
			}
		}
		if a%2 == 1 && b%2 == 1 { // 在该位1的次数都为奇数：奇数x奇数=奇数
			res = res | (1 << i)
		}
	}
	return res
}
