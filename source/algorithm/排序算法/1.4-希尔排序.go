package main

import "fmt"

func main() {
	//arr := []int{66, 33, 55, 22, 11, 99, 88, 77, 8}
	arr := []int{5, 2, 3, 1}
	shellSort(arr)
	fmt.Println(arr)
}

func shellSort(arr []int) {
	n := len(arr)
	for gap := n / 2; gap > 0; gap = gap / 2 {
		for i := gap; i < n; i++ {
			j := i
			cur := arr[i]
			for j-gap >= 0 && cur < arr[j-gap] {
				arr[j] = arr[j-gap]
				j = j - gap
			}
			arr[j] = cur
		}
	}
}
