package main

import "fmt"

func main() {
	arr := []int{66, 33, 55, 22, 11, 99, 88, 77}
	insertionSort(arr)
	fmt.Println(arr)
}

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		pos := i - 1 // 已经有序的最后一个下标
		cur := arr[i]
		for pos >= 0 && arr[pos] > cur {
			arr[pos+1] = arr[pos] // 后移
			pos--
		}
		arr[pos+1] = cur
	}
}
