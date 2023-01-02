package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	arrA := make([]int, n)
	arrB := make([]int, n)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		arrA[i] = a
		arrB[i] = b
	}
	res := getResult(arrA, arrB)
	fmt.Println(res)
}

func getResult(a, b []int) int {
	sum := 0
	res := 0
	for i := 0; i < len(a); i++ {
		sum = sum + b[i] - a[i]
		if sum > res {
			res = sum
		}
	}
	return res
}
