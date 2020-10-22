package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(getKth(10, 20, 5))
}

// leetcode1387_将整数按权重排序
func getKth(lo int, hi int, k int) int {
	arr := make([][2]int, 0)
	for i := lo; i <= hi; i++ {
		arr = append(arr, [2]int{i, getCount(i)})
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i][1] == arr[j][1] {
			return arr[i][0] < arr[j][0]
		}
		return arr[i][1] < arr[j][1]
	})
	return arr[k-1][0]
}

func getCount(i int) int {
	if i == 1 {
		return 0
	}
	if i%2 == 1 {
		return getCount(i*3+1) + 1
	}
	return getCount(i/2) + 1
}
