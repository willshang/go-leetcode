package main

import "sort"

func main() {

}

func splitPainting(segments [][]int) [][]int64 {
	res := make([][]int64, 0)
	m := make(map[int]int)
	for i := 0; i < len(segments); i++ {
		start := segments[i][0]
		end := segments[i][1]
		count := segments[i][2]
		m[start] = m[start] + count
		m[end] = m[end] - count
	}
	arr := make([][2]int, 0)
	for k, v := range m {
		arr = append(arr, [2]int{k, v})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})
	n := len(arr)
	for i := 1; i < n; i++ { // 前缀和
		arr[i][1] = arr[i][1] + arr[i-1][1]
	}
	for i := 0; i < n-1; i++ {
		if arr[i][1] > 0 { // 和大于0
			res = append(res, []int64{
				int64(arr[i][0]),
				int64(arr[i+1][0]),
				int64(arr[i][1]),
			})
		}
	}
	return res
}
