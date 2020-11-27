package main

import "fmt"

func main() {
	arr := []int{66, 33, 55, 22, 11, 99, 88, 77, 8}
	//arr := []int{5, 2, 3, 1}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	index := partition(arr, left, right)
	quickSort(arr, left, index-1)
	quickSort(arr, index+1, right)
}

func partition(arr []int, left, right int) int {
	baseValue := arr[left] // 基准值
	mark := left
	for i := left + 1; i <= right; i++ {
		if arr[i] < baseValue {
			mark++
			arr[mark], arr[i] = arr[i], arr[mark]
		}
	}
	arr[left], arr[mark] = arr[mark], arr[left]
	return mark
}
