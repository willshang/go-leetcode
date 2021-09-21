package main

import "sort"

func main() {

}

func maxTaxiEarnings(n int, rides [][]int) int64 {
	m := len(rides)
	arr := make([][]int64, 0)
	for i := 0; i < m; i++ {
		a, b, c := rides[i][0], rides[i][1], rides[i][2]
		arr = append(arr, []int64{int64(a), int64(b), int64(b - a + c)})
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i][1] == arr[j][1] {
			return arr[i][0] < arr[j][0]
		}
		return arr[i][1] < arr[j][1]
	})
	for i := 1; i < m; i++ {
		target := sort.Search(i, func(j int) bool {
			return arr[j][1] > arr[i][0]
		})
		if target == 0 {
			arr[i][2] = max(arr[i][2], arr[i-1][2])
		} else {
			arr[i][2] = max(arr[i][2]+arr[target-1][2], arr[i-1][2])
		}
	}
	return arr[m-1][2]
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
