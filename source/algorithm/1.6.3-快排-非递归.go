package main

import "fmt"

func main() {
	arr := []int{66, 33, 55, 22, 11, 99, 88, 77, 8}
	//arr := []int{5, 2, 3, 1}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quickSort(arr []int, left, right int) {
	queue := make([][2]int, 0)
	queue = append(queue, [2]int{left, right})
	for len(queue) > 0 {
		l, r := queue[0][0], queue[0][1]
		queue = queue[1:]
		index := partition(arr, l, r)
		if l < index-1 {
			queue = append(queue, [2]int{l, index - 1})
		}
		if index+1 < r {
			queue = append(queue, [2]int{index + 1, r})
		}
	}
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
