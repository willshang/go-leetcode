package main

import "fmt"

func main() {
	fmt.Println(findLatestStep([]int{3, 5, 1, 2, 4}, 1))
	fmt.Println(findLatestStep([]int{3, 2, 5, 6, 10, 8, 9, 4, 1, 7}, 3))
}

// leetcode1562_查找大小为M的最新分组
func findLatestStep(arr []int, m int) int {
	res := -1
	temp := make([]int, len(arr)+2)
	M := make(map[int]bool)
	for i := 0; i < len(arr); i++ {
		index := arr[i]
		temp[index] = 1
		left := temp[index-1]
		right := temp[index+1]
		total := 1 + left + right
		temp[index] = total
		temp[index-left] = total
		temp[index+right] = total
		for k := range M {
			if index-left <= k && k < index {
				delete(M, k)
			}
			if index < k && k <= index+right {
				delete(M, k)
			}
		}
		if total == m {
			M[index] = true
		}
		if len(M) > 0 {
			res = i + 1
		}
	}
	return res
}
