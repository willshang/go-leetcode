package main

import "fmt"

func main() {
	fmt.Println(findLatestStep([]int{3, 5, 1, 2, 4}, 1))
	fmt.Println(findLatestStep([]int{3, 2, 5, 6, 10, 8, 9, 4, 1, 7}, 3))
}

func findLatestStep(arr []int, m int) int {
	if len(arr) == m {
		return len(arr)
	}
	res := -1
	temp := make([]int, len(arr)+2)
	for i := 0; i < len(arr); i++ {
		index := arr[i]
		total := 1 + temp[index-1] + temp[index+1]
		// 上一步一定存在m
		if temp[index-1] == m || temp[index+1] == m {
			res = i
		}
		temp[index-temp[index-1]] = total
		temp[index+temp[index+1]] = total
	}
	return res
}
