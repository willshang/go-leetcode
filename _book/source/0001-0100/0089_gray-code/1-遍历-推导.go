package main

import "fmt"

func main() {
	fmt.Println(grayCode(5))
}

func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}
	res := []int{0, 1}
	for i := 1; i < n; i++ {
		temp := make([]int, 0)
		value := 1 << i
		for j := len(res) - 1; j >= 0; j-- {
			// 10 1 11
			// 10 0 10
			// 100 10 110
			// 100 11 111
			// 100 1 101
			// 100 0 100
			// fmt.Printf("%b %b %b\n", value, res[j], res[j]^value)

			// temp = append(temp, res[j]|value)
			// temp = append(temp, res[j]^value)
			temp = append(temp, res[j]+value)
		}
		res = append(res, temp...)
	}
	return res
}
