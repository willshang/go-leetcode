package main

import (
	"fmt"
)

func main() {
	//fmt.Println(getLeastNumbers([]int{6, 3, 2, 1, 4, 5, 9, 7, 8}, 2))
	//fmt.Println(getLeastNumbers([]int{0, 0, 1, 2, 4, 2, 2, 3, 1, 4}, 8))
	//fmt.Println(getLeastNumbers([]int{3, 2, 1}, 2))

	arr := []int{0, 1, 1, 1, 4, 5, 3, 7, 7, 8, 10, 2, 7, 8, 0, 5, 2, 16, 12, 1, 19, 15, 5, 18, 2, 2, 22, 15, 8, 22, 17, 6, 22, 6, 22, 26, 32,
		8, 10, 11, 2, 26, 9, 12, 9, 7, 28, 33, 20, 7, 2, 17, 44, 3, 52, 27, 2, 23, 19, 56, 56, 58, 36, 31, 1, 19, 19,
		6, 65, 49, 27, 63, 29, 1, 69, 47, 56, 61, 40, 43, 10, 71, 60, 66, 42, 44, 10, 12, 83, 69, 73, 2, 65, 93, 92, 47, 35, 39, 13, 75}
	fmt.Println(getLeastNumbers(arr, 75))
}

// 剑指offer_面试题40_最小的k个数
func getLeastNumbers(arr []int, k int) []int {
	if len(arr) == 0 || k == 0 {
		return nil
	}
	left := 0
	right := len(arr) - 1
	for {
		index := partition(arr, left, right)
		if index+1 == k {
			break
		} else if index+1 < k {
			left = index + 1
		} else {
			right = index - 1
		}
	}
	return arr[:k]
}

func partition(arr []int, left, right int) int {
	if left == right {
		return left
	}
	value := arr[left]
	i := left
	j := right
	for {
		for arr[j] >= value && i < j {
			j--
		}
		for arr[i] <= value && i < j {
			i++
		}
		if i >= j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[left] = arr[i]
	arr[i] = value
	return i
}
