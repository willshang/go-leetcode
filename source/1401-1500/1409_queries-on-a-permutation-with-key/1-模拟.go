package main

import "fmt"

func main() {
	fmt.Println(processQueries([]int{3, 1, 2, 1}, 5))
}

// leetcode1409_查询带键的排列
func processQueries(queries []int, m int) []int {
	n := len(queries)
	res := make([]int, n)
	arr := make([]int, m)
	for i := 0; i < m; i++ {
		arr[i] = i + 1
	}
	for i := 0; i < n; i++ {
		index := 0
		for j := 0; j < m; j++ {
			if arr[j] == queries[i] {
				index = j
				break
			}
		}
		res[i] = index
		arr = append(append([]int{arr[index]}, arr[:index]...), arr[index+1:]...)
	}
	return res
}
