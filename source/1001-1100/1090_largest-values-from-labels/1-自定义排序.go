package main

import "sort"

func main() {

}

// leetcode1090_受标签影响的最大值
func largestValsFromLabels(values []int, labels []int, num_wanted int, use_limit int) int {
	arr := make([][2]int, 0)
	for i := 0; i < len(values); i++ {
		arr = append(arr, [2]int{values[i], labels[i]})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] > arr[j][0]
	})
	res := 0
	m := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		if m[arr[i][1]] < use_limit {
			res = res + arr[i][0]
			m[arr[i][1]]++
			num_wanted--
		}
		if num_wanted <= 0 {
			break
		}
	}
	return res
}
