package main

import "fmt"

func main() {
	fmt.Println(sumOddLengthSubarrays([]int{1, 4, 2, 5, 3}))
	fmt.Println(sumOddLengthSubarrays([]int{1, 2}))
	fmt.Println(sumOddLengthSubarrays([]int{10, 11, 12}))
}

func sumOddLengthSubarrays(arr []int) int {
	res := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		left := i + 1
		right := n - i
		leftEven := (left + 1) / 2
		rightEven := (right + 1) / 2
		leftOdd := left / 2
		rightOdd := right / 2
		res = res + arr[i]*(leftEven*rightEven+leftOdd*rightOdd)
	}
	return res
}
