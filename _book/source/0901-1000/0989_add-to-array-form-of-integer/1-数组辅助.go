package main

import "fmt"

func main() {
	fmt.Println(addToArrayForm([]int{1, 2, 0, 0}, 34))
}

func addToArrayForm(A []int, K int) []int {
	B := make([]int, 0)
	for K > 0 {
		B = append([]int{K % 10}, B...)
		K = K / 10
	}
	length := len(A)
	if len(B) > len(A) {
		length = len(B)
	}
	res := make([]int, length)
	flag := 0
	i := len(A) - 1
	j := len(B) - 1
	count := 0
	for i >= 0 && j >= 0 {
		sum := A[i] + B[j] + flag
		if sum >= 10 {
			sum = sum - 10
			flag = 1
		} else {
			flag = 0
		}
		res[length-1-count] = sum
		i--
		j--
		count++
	}
	for i >= 0 {
		sum := A[i] + flag
		if sum >= 10 {
			sum = sum - 10
			flag = 1
		} else {
			flag = 0
		}
		res[length-1-count] = sum
		i--
		count++
	}
	for j >= 0 {
		sum := B[j] + flag
		if sum >= 10 {
			sum = sum - 10
			flag = 1
		} else {
			flag = 0
		}
		res[length-1-count] = sum
		j--
		count++
	}
	if flag == 1 {
		return append([]int{1}, res...)
	}
	return res
}
