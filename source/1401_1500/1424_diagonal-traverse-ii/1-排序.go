package main

import "sort"

func main() {

}

func findDiagonalOrder(nums [][]int) []int {
	arr := make([][2]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			arr = append(arr, [2]int{i, j})
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		a := arr[i][0] + arr[i][1]
		b := arr[j][0] + arr[j][1]
		if a == b {
			return arr[i][1] < arr[j][1]
		}
		return a < b
	})
	res := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		a, b := arr[i][0], arr[i][1]
		res = append(res, nums[a][b])
	}
	return res
}
