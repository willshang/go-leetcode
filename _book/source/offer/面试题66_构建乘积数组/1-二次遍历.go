package main

import "fmt"

func main() {
	fmt.Println(constructArr([]int{1, 2, 3, 4, 5}))
}

// 剑指offer_面试题66_构建乘积数组
func constructArr(a []int) []int {
	res := make([]int, len(a))
	if len(a) == 0 {
		return res
	}
	res[0] = 1
	for i := 1; i < len(res); i++ {
		res[i] = res[i-1] * a[i-1]
	}
	temp := 1
	for i := len(res) - 2; i >= 0; i-- {
		res[i] = res[i] * a[i+1] * temp
		temp = temp * a[i+1]
	}
	return res
}
