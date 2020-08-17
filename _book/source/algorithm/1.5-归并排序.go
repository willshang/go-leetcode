package main

import "fmt"

func main() {
	arr := []int{66, 33, 55, 22, 11, 99, 88, 77}
	fmt.Println(mergeSort(arr))
	fmt.Println(arr)
}

func mergeSort(arr []int) []int {
	n := len(arr)
	if n < 2 {
		return arr
	}
	return merge(mergeSort(arr[:len(arr)/2]), mergeSort(arr[len(arr)/2:]))
}

func merge(left, right []int) []int {
	res := make([]int, 0)
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			res = append(res, left[0])
			left = left[1:]
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}
	if len(left) > 0 {
		res = append(res, left...)
	}
	if len(right) > 0 {
		res = append(res, right...)
	}
	return res
}
