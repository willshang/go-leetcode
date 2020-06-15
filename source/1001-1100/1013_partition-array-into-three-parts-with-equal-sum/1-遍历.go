package main

import "fmt"

func main() {
	fmt.Println(canThreePartsEqualSum([]int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4}))
}

// leetcode1013_将数组分成和相等的三个部分
func canThreePartsEqualSum(A []int) bool {
	length := len(A)
	if length < 3 {
		return false
	}
	sum := 0
	for i := 0; i < length; i++ {
		sum = sum + A[i]
	}
	if sum%3 != 0 {
		return false
	}
	target := sum / 3
	count := 0
	temp := 0
	for i := 0; i < len(A); i++ {
		temp = temp + A[i]
		if temp == target {
			temp = 0
			count++
		}
	}
	if count >= 3 {
		return true
	}
	return false
}
