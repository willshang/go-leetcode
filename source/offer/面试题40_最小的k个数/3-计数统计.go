package main

import (
	"fmt"
)

func main() {
	fmt.Println(getLeastNumbers([]int{3, 2, 1}, 1))
}

func getLeastNumbers(arr []int, k int) []int {
	if len(arr) == 0 || k == 0 {
		return nil
	}
	a := make([]int, 10001)
	for _, v := range arr {
		a[v]++
	}
	res := make([]int, 0)
	for key, value := range a {
		if value > 0 {
			for i := 0; i < value; i++ {
				res = append(res, key)
				k--
				if k <= 0 {
					return res
				}
			}
		}
	}
	return res
}
