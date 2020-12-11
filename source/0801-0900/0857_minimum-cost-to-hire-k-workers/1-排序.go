package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(mincostToHireWorkers([]int{10, 20, 5}, []int{70, 50, 30}, 2))
}

func mincostToHireWorkers(quality []int, wage []int, K int) float64 {
	arr := make([][2]int, 0)
	for i := 0; i < len(quality); i++ {
		arr = append(arr, [2]int{quality[i], wage[i]})
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i][1]/arr[i][0] == arr[j][1]/arr[j][0] {
			return arr[i][0] < arr[j][0]
		}
		return arr[i][1]/arr[i][0] < arr[j][1]/arr[j][0]
	})
	fmt.Println(arr)
	res := float64(0)
	res = res + float64(arr[0][1])
	per := float64(arr[0][1]) / float64(arr[0][0])
	for i := 1; i < K; i++ {
		res = res + per*float64(arr[i][0])
	}
	return res
}
