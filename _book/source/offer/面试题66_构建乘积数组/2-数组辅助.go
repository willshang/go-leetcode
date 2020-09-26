package main

import "fmt"

func main() {
	fmt.Println(constructArr([]int{1, 2, 3, 4, 5}))
}

func constructArr(a []int) []int {
	res := make([]int, len(a))
	if len(a) == 0 {
		return res
	}
	left := make([]int, len(a))
	left[0] = 1
	right := make([]int, len(a))
	right[len(a)-1] = 1
	for i := 1; i < len(a); i++ {
		left[i] = left[i-1] * a[i-1]
	}
	for i := len(a) - 2; i >= 0; i-- {
		right[i] = right[i+1] * a[i+1]
	}
	for i := 0; i < len(a); i++ {
		res[i] = left[i] * right[i]
	}
	return res
}
