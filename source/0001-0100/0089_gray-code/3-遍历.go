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
		value := 1 << i
		for j := len(res) - 1; j >= 0; j-- {
			res = append(res, res[j]+value)
		}
	}
	return res
}
