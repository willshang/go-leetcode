package main

import "fmt"

func main() {
	fmt.Println(pancakeSort([]int{3, 2, 4, 1}))
}

// leetcode969_煎饼排序
func pancakeSort(arr []int) []int {
	res := make([]int, 0)
	n := len(arr) - 1
	for n >= 0 {
		maxValue := arr[0]
		index := 0
		for i := 1; i <= n; i++ {
			if arr[i] > maxValue {
				maxValue = arr[i]
				index = i
			}
		}
		if index == n {
			n--
			continue
		}
		if index != 0 {
			res = append(res, index+1)
			reverse(arr, 0, index)
		}
		res = append(res, n+1)
		reverse(arr, 0, n)
		n--
	}
	return res
}

func reverse(arr []int, left, right int) {
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}
