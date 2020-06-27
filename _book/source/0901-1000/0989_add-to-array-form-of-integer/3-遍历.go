package main

import "fmt"

func main() {
	fmt.Println(addToArrayForm([]int{1, 2, 0, 0}, 34))
}

func addToArrayForm(A []int, K int) []int {
	i := len(A) - 1
	res := make([]int, 0)
	for i >= 0 || K > 0 {
		if i >= 0 {
			K = K + A[i]
		}
		res = append(res, K%10)
		K = K / 10
		i--
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}
