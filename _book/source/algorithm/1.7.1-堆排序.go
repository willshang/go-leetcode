package main

import "fmt"

func main() {
	arr := []int{66, 33, 55, 22, 11, 99, 88, 77}
	heapSort(arr)
	fmt.Println(arr)
}

func heapSort(arr []int) {
	buildHeap(arr, len(arr))
	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapModify(arr, 0, i-1)
	}
}

func buildHeap(arr []int, length int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		heapModify(arr, i, length-1)
	}
}

func heapModify(arr []int, start, end int) {
	temp := arr[start]
	for left := 2*start + 1; left <= end; left = 2*left + 1 {
		if left < end && arr[left] < arr[left+1] {
			left++
		}
		if arr[start] < arr[left] {
			arr[start] = arr[left]
			start = left
		}
		arr[start] = temp
	}
}
