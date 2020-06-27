package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(arrayRankTransform([]int{40, 10, 20, 30}))
}

// leetcode1331_数组序号转换
func arrayRankTransform(arr []int) []int {
	temp := make([]int, len(arr))
	copy(temp, arr)
	sort.Ints(temp)
	m := make(map[int]int)
	count := 1
	for i := 0; i < len(temp); i++ {
		if m[temp[i]] > 0 {
			continue
		}
		m[temp[i]] = count
		count++
	}
	res := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		res[i] = m[arr[i]]
	}
	return res
}
