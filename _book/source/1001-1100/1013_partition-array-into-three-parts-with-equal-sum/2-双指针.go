package main

import "fmt"

func main() {
	//fmt.Println(canThreePartsEqualSum([]int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4}))
	fmt.Println(canThreePartsEqualSum([]int{1, -1, 1, -1}))
}

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
	left, right := 1, len(A)-2
	leftValue, rightValue := A[0], A[len(A)-1]
	for left < right {
		for left < right && leftValue != target {
			leftValue = leftValue + A[left]
			left++
		}
		for left < right && rightValue != target {
			rightValue = rightValue + A[right]
			right--
		}
		if leftValue == target && rightValue == target {
			return true
		}
	}
	return false
}
