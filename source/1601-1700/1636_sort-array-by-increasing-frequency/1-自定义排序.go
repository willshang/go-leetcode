package main

import "sort"

func main() {

}

// leetcode1636_按照频率将数组升序排序
func frequencySort(nums []int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	arr := make([][2]int, 0)
	for k, v := range m {
		arr = append(arr, [2]int{k, v})
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i][1] == arr[j][1] {
			return arr[i][0] > arr[j][0]
		}
		return arr[i][1] < arr[j][1]
	})
	res := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < arr[i][1]; j++ {
			res = append(res, arr[i][0])
		}
	}
	return res
}
