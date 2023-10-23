package main

import "fmt"

func main() {
	fmt.Println(processQueries([]int{3, 1, 2, 1}, 5))
}

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
				// 交换位置
				for k := 0; k < index; k++ {
					arr[k], arr[index] = arr[index], arr[k]
				}
				break
			}
		}
		res[i] = index
	}
	return res
}
