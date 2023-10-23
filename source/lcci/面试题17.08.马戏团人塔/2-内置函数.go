package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(bestSeqAtIndex([]int{65, 70, 56, 75, 60, 68}, []int{100, 150, 0, 190, 95, 110}))
}

func bestSeqAtIndex(height []int, weight []int) int {
	arr := make([][2]int, 0)
	for i := 0; i < len(height); i++ {
		arr = append(arr, [2]int{height[i], weight[i]})
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i][0] == arr[j][0] {
			return arr[i][1] < arr[j][1]
		}
		return arr[i][0] > arr[j][0]
	})
	res := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		index := sort.Search(len(res), func(j int) bool {
			return arr[res[j]][0] <= arr[i][0] || arr[res[j]][1] <= arr[i][1]
		})
		if index == len(res) {
			res = append(res, i)
		} else {
			res[index] = i
		}
	}
	return len(res)
}
