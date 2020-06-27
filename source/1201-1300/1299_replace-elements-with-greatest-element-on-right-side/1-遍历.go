package main

import "fmt"

func main() {
	fmt.Println(replaceElements([]int{17, 18, 5, 4, 6, 1}))
}

// leetcode1299_将每个元素替换为右侧最大元素
func replaceElements(arr []int) []int {
	max := -1
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] > max {
			arr[i], max = max, arr[i]
		} else {
			arr[i] = max
		}
	}
	return arr
}
