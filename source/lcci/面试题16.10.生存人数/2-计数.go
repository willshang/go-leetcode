package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxAliveYear([]int{1900, 1901, 1950}, []int{1948, 1951, 2000}))
}

// 程序员面试金典16.10_生存人数
func maxAliveYear(birth []int, death []int) int {
	arr := make([]int, 102)
	for i := 0; i < len(birth); i++ {
		arr[birth[i]-1900]++
		arr[death[i]-1900+1]--
	}
	max := 0
	sum := 0
	res := 0
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
		if sum > max {
			max = sum
			res = i + 1900
		}
	}
	return res
}
